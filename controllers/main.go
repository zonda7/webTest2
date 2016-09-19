package controllers

import (
	//"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"webTest/libs"
	"webTest/models"
)

type MainController struct {
	BaseController
}

//首页
func (this *MainController) Index() {

	this.TplName = "layout/layout.html"

}

func (this *MainController) Logs() {

	paths := libs.GetFilelist(beego.AppConfig.String("file.path"))

	this.Data["paths"] = paths

	this.display()
}

//运行日志详细
func (this *MainController) ViewLog() {

	file := strings.TrimSpace(this.GetString("path"))

	this.Data["str"] = libs.Read(file)

	this.display()
}

//登录
func (this *MainController) Login() {

	if this.userId > 0 {
		this.redirect("/")
	}

	beego.ReadFromRequest(&this.Controller)

	if this.isPost() {
		flash := beego.NewFlash()
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		//remember := this.GetString("remember")
		worryCount := 0
		if username != "" && password != "" {

			user, err := models.UserGetByName(username)
			errorMsg := ""

			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {

				//models.UserUpdate(user)

				if user != nil {
					worryCount = user.WorryCount

					if worryCount > 5 {
						errorMsg = "帐号已销定"
					} else {
						user.WorryCount = worryCount + 1
						models.UserUpdate(user)
						errorMsg = "帐号或密码错误"
					}

				} else {
					errorMsg = "帐号或密码错误"
				}

			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {

				user.WorryCount = 0
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				models.UserUpdate(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))

				this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 3000)

				/*
					if remember == "yes" {
						this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 60)

					} else {
						this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
					}*/

				if user.LoginCount == 0 { //第一次登录修改密码
					this.redirect(beego.URLFor("MainController.Profile"))
				}

				this.redirect(beego.URLFor("MainController.Logs"))

			}

			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("MainController.Login"))

		}

	}

	this.TplName = "main/login.html"

}

//获取用户信息
func (this *MainController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.UserGetList(page, this.pageSize)

	this.Data["pageTitle"] = "分组列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("MainController.List"), true).ToString()
	this.display()
}

//用户添加
func (this *MainController) Add() {

	if this.isPost() {

		user := new(models.User)
		user.UserName = strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		user.Email = strings.TrimSpace(this.GetString("email"))
		user.Role, _ = this.GetInt("role")

		user.Salt = string(utils.RandomCreateBytes(10))
		user.Password = libs.Md5([]byte(password + user.Salt))
		println(user.Password)
		_, err := models.UserAdd(user)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)

	}
	this.Data["pageTitle"] = "用户添加"
	this.display()

}

//修改个人信息
func (this *MainController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user, err := models.UserGetById(this.userId)
	if err != nil {
		beego.Error(err)
	}
	if this.isPost() {
		flash := beego.NewFlash()
		user.Email = this.GetString("email")
		user.Update()
		password1 := this.GetString("password1")

		password2 := this.GetString("password2")

		if password1 != "" {
			if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else if password2 != password1 {
				flash.Error("两次输入的密码不一致")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else {

				user.Salt = string(utils.RandomCreateBytes(10))
				user.Password = libs.Md5([]byte(password1 + user.Salt))

				logincount, _ := this.GetInt("login_count") //是否首次登录
				if logincount == 0 {
					user.LoginCount = logincount + 1
				}
				user.Update()
			}
		} else {
			flash.Error("请输入密码")
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor(".Profile"))

		}

		flash.Success("修改成功！")
		flash.Store(&this.Controller)
		this.redirect(beego.URLFor(".Profile"))

	}
	this.Data["pageTitle"] = "个人信息"
	this.Data["user"] = user
	this.display()

}

//用户处理
func (this *MainController) Batch() {
	action := this.GetString("action")
	ids := this.GetStrings("ids")
	if len(ids) < 1 {
		this.ajaxMsg("请选择要操作的项目", MSG_ERR)
	}

	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		if id < 1 {
			continue
		}
		switch action {
		case "delete":
			models.UserDelById(id)
		case "password":

			user, err := models.UserGetById(id)

			if err != nil {
				beego.Error(err)
			}
			println(id)

			user.Salt = string(utils.RandomCreateBytes(10))
			user.Password = libs.Md5([]byte("123456" + user.Salt))
			user.Update()

		}
	}

	this.ajaxMsg("", MSG_OK)

}

//用户退出
func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("MainController.Login"))
}

//重启
func (this *MainController) Reboot() {

	if this.isPost() {

		cmd := exec.Command("reboot")
		cmd.Run()

	}

	remoteis := 0
	id := models.SystemGetList()

	system, _ := models.SystemGetById(id)

	if system != nil {
		remoteis = system.RemoteIs
	}

	this.Data["remoteis"] = remoteis
	this.Data["pageTitle"] = "系统设置"
	this.display()
}

//远程登录打开
func (this *MainController) RemoteOpen() {

	id := models.SystemGetList()

	system := new(models.System)
	system.RemoteIp = ""
	system.RemoteMac = ""
	system.RemoteIs = 1

	if id == 0 {

		models.SystemAdd(system)

	} else {
		system.Id = id
		system.Update()
	}

	cmd := exec.Command("/bin/bash", "-c", "iptables -I INPUT -p tcp --dport 22 -j DROP")
	cmd.Run()

	refer := this.Ctx.Request.Referer()
	if refer == "" {
		refer = beego.URLFor("/reboot")
	}
	this.redirect(refer)

}

//远程关闭（默认关闭远程登录）
func (this *MainController) RemoteClose() {

	id := models.SystemGetList()
	system := new(models.System)
	system.RemoteIp = ""
	system.RemoteMac = ""

	if id == 0 { //如果打开第一次打开程序关闭程序
		system.RemoteIs = 0
		models.SystemAdd(system)

	} else {
		system.Id = id
		system.RemoteIs = 0
		system.Update()
	}

	cmd := exec.Command("/bin/bash", "-c", "iptables -I INPUT -p tcp --dport 22 -j ACCEPT")
	cmd.Run()

	refer := this.Ctx.Request.Referer()
	if refer == "" {
		refer = beego.URLFor("/reboot")
	}
	this.redirect(refer)

}
