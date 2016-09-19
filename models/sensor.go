package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Sensor struct {
	Id           int
	SensorName   string
	ChannelName  string
	DecimalPlace int
	GroupNumber  int
	LimitField   string
	ItemData     string
	TargetField  string
	NameTable    string
}

func (s *Sensor) TableName() string {
	return TableName("sensor_setting")
}

func (s *Sensor) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

//新增
func SensorAdd(sensor *Sensor) (int64, error) {
	if sensor.SensorName == "" {
		return 0, fmt.Errorf("SensorName字段不能为空")
	}

	return orm.NewOrm().Insert(sensor)
}

//查询
func SensorGetList(page, pageSize int, filters ...interface{}) ([]*Sensor, int64) {
	offset := (page - 1) * pageSize

	sensors := make([]*Sensor, 0)

	query := orm.NewOrm().QueryTable(TableName("sensor_setting"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&sensors)

	return sensors, total

}

//更新状态
func TaskResetGroupId(groupNumber int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("sensor_setting")).Filter("group_number", groupNumber).Update(orm.Params{
		"group_number": 0,
	})
}

//根据ID查询
func SersorGetById(id int) (*Sensor, error) {
	sensor := &Sensor{
		Id: id,
	}
	err := orm.NewOrm().Read(sensor)
	if err != nil {
		return nil, err
	}
	return sensor, nil

}

//删除
func SersorDel(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("sensor_setting")).Filter("id", id).Delete()
	return err
}
