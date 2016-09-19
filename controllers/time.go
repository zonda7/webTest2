package controllers

import (
	"github.com/astaxie/beego"
	"webTest/libs"
	"webTest/models"
)

type TimeController struct {
	BaseController
}

func (this *TimeController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.TimeGetList(page, this.pageSize)

	this.Data["pageTitle"] = "实时数据列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("TimeController.List"), true).ToString()
	this.display()
}
