package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L. -lopcIp
#include "py_api.h"
*/
import "C"
import (
	"github.com/astaxie/beego"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"webTest/controllers"
	"webTest/models"
)

func main() {

	//https服务
	go (func() {

		var rpURL *url.URL
		var err error

		if rpURL, err = url.Parse("http://" + beego.AppConfig.String("httpaddr") + ":" + beego.AppConfig.String("httpport")); err != nil {
			log.Fatal(err)
		}

		reverse := httputil.NewSingleHostReverseProxy(rpURL)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			reverse.ServeHTTP(w, r)
		})

		http.ListenAndServeTLS(":8081", "server.crt",
			"server.key", nil)

	})()

	//打开ipv4服务
	//C.PyInit()

	models.Init()

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/list", &controllers.MainController{}, "*:List")
	beego.Router("/batch", &controllers.MainController{}, "*:Batch")
	beego.Router("/add", &controllers.MainController{}, "*:Add")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/help", &controllers.MainController{}, "*:Index")
	beego.Router("/logs", &controllers.MainController{}, "*:Logs")
	beego.Router("/viewlog", &controllers.MainController{}, "*:ViewLog")

	beego.Router("/remoteclose", &controllers.MainController{}, "*:RemoteClose")
	beego.Router("/remoteopen", &controllers.MainController{}, "*:RemoteOpen")
	beego.Router("/reboot", &controllers.MainController{}, "*:Reboot")

	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.SensorController{})
	beego.AutoRouter(&controllers.TimeController{})
	beego.AutoRouter(&controllers.NetController{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()

}
