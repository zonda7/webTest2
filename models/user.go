package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	UserName   string
	Password   string
	Salt       string
	Email      string
	LastLogin  int64
	LastIp     string
	Status     int
	LoginCount int
	WorryCount int
	Role       int
}

func (u *User) TableName() string {
	return TableName("user")
}

//添加用户
func UserAdd(user *User) (int64, error) {
	return orm.NewOrm().Insert(user)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

//使用用户名获取用户
func UserGetByName(userName string) (*User, error) {
	u := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("user_name", userName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserGetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize

	users := make([]*User, 0)

	query := orm.NewOrm().QueryTable(TableName("user"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&users)

	return users, total
}

//使用Id获取用户
func UserGetById(id int) (*User, error) {
	u := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//更新用户
func UserUpdate(user *User, fields ...string) error {
	_, err := orm.NewOrm().Update(user, fields...)
	return err
}

//删除用户
func UserDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).Delete()
	return err
}
