package models

import (
	"github.com/astaxie/beego/orm"
)

type System struct {
	Id        int
	RemoteIp  string
	RemoteMac string
	RemoteIs  int
}

func (s *System) TableName() string {
	return TableName("system_setting")
}

func (s *System) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

func SystemAdd(system *System) (int64, error) {

	return orm.NewOrm().Insert(system)
}

func SystemGetList() int {

	id := 0

	list := make([]*System, 0)

	query := orm.NewOrm().QueryTable(TableName("system_setting"))

	query.Limit(1, 0).All(&list)

	for _, value := range list {

		id = value.Id
	}

	return id
}

func SystemGetById(id int) (*System, error) {
	system := &System{
		Id: id,
	}

	err := orm.NewOrm().Read(system)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func SystemDel(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("system_setting")).Filter("id", id).Delete()
	return err
}
