package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	//"net/url"
)

func Init() {

	//orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "gateway_data_fetch.db")

	orm.RegisterModel(new(User), new(TaskGroup), new(Time), new(Sensor), new(Net), new(TaskLog), new(System))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

func TableName(name string) string {
	return name
}

func aa(id int) int {
	return id
}
