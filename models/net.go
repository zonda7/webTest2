package models

import (
	"github.com/astaxie/beego/orm"
)

type Net struct {
	Id               int
	UdpServerIp      string
	UdpServerPort    int
	DbComputerIp     string
	DbDatabaseName   string
	DbUserName       string
	DbUserPassword   string
	DbFetchInterval  int
	OpcComputerIp    string
	OpcDatabaseName  string
	OpcUserName      string
	OpcUserPassword  string
	OpcFetchInterval int
	TargetUdpIp      string
	TargetUdpPort    int
	TypeFlag         int
}

//IP结构体
type Ipv4 struct {
	DwLocalIp  string
	DwNetMask  string
	DwGateWay  string
	DwFirstDns string
	DwSecDns   string
}

//IP结构体
type IpInfo struct {
	DwLocalIp  uint32
	DwNetMask  uint32
	DwGateWay  uint32
	DwFirstDns uint32
	DwSecDns   uint32
}

func (n *Net) TableName() string {

	return TableName("config_setting")
}

//更新
func (t *Net) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func NetAdd(obj *Net) (int64, error) {

	return orm.NewOrm().Insert(obj)
}

func NetGetById(id int) (*Net, error) {
	net := &Net{
		Id: id,
	}

	err := orm.NewOrm().Read(net)
	if err != nil {
		return nil, err
	}
	return net, nil
}

func NetGetList() int {

	id := 0

	list := make([]*Net, 0)

	query := orm.NewOrm().QueryTable(TableName("config_setting"))

	query.Limit(1, 0).All(&list)

	for _, value := range list {

		id = value.Id
	}

	return id
}
