package models

import (
	"github.com/astaxie/beego/orm"
)

type Time struct {
	Id         int
	Content    string
	UpdateTime string
}

func (t *Time) TableName() string {

	return TableName("realtime_message")
}

func TimeGetList(page, pageSize int) ([]*Time, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Time, 0)

	query := orm.NewOrm().QueryTable(TableName("realtime_message"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
