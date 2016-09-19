package controllers

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L.. -lopcIp
#include "py_api.h"
*/
import "C"
import (
	"strings"
	"webTest/libs"
	"webTest/models"
)

type NetController struct {
	BaseController
}

func (this *NetController) Edit() {

	id := models.NetGetList()

	if this.isPost() {

		net := new(models.Net)

		net.UdpServerIp = strings.TrimSpace(this.GetString("udp_server_ip"))
		net.UdpServerPort, _ = this.GetInt("udp_server_port")
		net.DbComputerIp = strings.TrimSpace(this.GetString("db_computer_ip"))
		net.DbDatabaseName = strings.TrimSpace(this.GetString("db_database_name"))
		net.DbUserName = strings.TrimSpace(this.GetString("db_user_name"))
		net.DbUserPassword = strings.TrimSpace(this.GetString("db_user_password"))
		net.DbFetchInterval, _ = this.GetInt("db_fetch_interval")
		net.OpcComputerIp = strings.TrimSpace(this.GetString("opc_computer_ip"))
		net.OpcDatabaseName = strings.TrimSpace(this.GetString("opc_database_name"))
		net.OpcUserName = strings.TrimSpace(this.GetString("opc_user_name"))
		net.OpcUserPassword = strings.TrimSpace(this.GetString("opc_user_password"))
		net.OpcFetchInterval, _ = this.GetInt("opc_fetch_interval")
		net.TargetUdpIp = strings.TrimSpace(this.GetString("target_udp_ip"))
		net.TargetUdpPort, _ = this.GetInt("target_udp_port")
		net.TypeFlag, _ = this.GetInt("type_flag")

		if id > 0 {
			net.Id = id
			net.Update()
		} else {
			models.NetAdd(net)
		}

		this.ajaxMsg("修改成功！", MSG_OK)

	}

	if id == 0 {
		net := new(models.Net)

		ipv4 := GetIPv4()

		net.UdpServerIp = ipv4.DwLocalIp
		net.UdpServerPort = 9090
		net.DbComputerIp = ""
		net.DbDatabaseName = ""
		net.DbUserName = ""
		net.DbUserPassword = ""
		net.DbFetchInterval = 10
		net.OpcComputerIp = ""
		net.OpcDatabaseName = ""
		net.OpcUserName = ""
		net.OpcUserPassword = ""
		net.OpcFetchInterval = 10
		net.TargetUdpIp = ""
		net.TargetUdpPort = 9090
		net.TypeFlag = 0

		this.Data["net"] = net
	} else {
		net, _ := models.NetGetById(id)

		this.Data["net"] = net

	}

	this.Data["pageTitle"] = "网络设置"
	this.display()

}

func (this *NetController) SetIP() {
	ipv4 := GetIPv4()

	if this.isPost() {

		dwLocalIp := libs.Ip2long(strings.TrimSpace(this.GetString("dw_localIp")))
		dwNetMask := libs.Ip2long(strings.TrimSpace(this.GetString("dw_netMask")))
		dwGateWay := libs.Ip2long(strings.TrimSpace(this.GetString("dw_gateWay")))
		dwFirstDns := libs.Ip2long(strings.TrimSpace(this.GetString("dw_firstDns")))
		dwSecDns := libs.Ip2long(strings.TrimSpace(this.GetString("dw_secDns")))

		ip := &C.struct_ipv4_t{}
		ip.dwLocalIp = C.ulong(dwLocalIp)
		ip.dwNetMask = C.ulong(dwNetMask)
		ip.dwGateWay = C.ulong(dwGateWay)
		ip.dwFirstDns = C.ulong(dwFirstDns)
		ip.dwSecDns = C.ulong(dwSecDns)

		i := 0
		//i := C.py_set_ipv4(ip)

		if i == 1 {
			this.ajaxMsg("", MSG_OK)
		} else {
			this.ajaxMsg("修改失败！", MSG_ERR)
		}

	}

	this.Data["pageTitle"] = "IP配置"
	this.Data["ipv4"] = ipv4
	this.display()

}

//调用c获取IPv4
func GetIPv4() *models.Ipv4 {

	ipv4 := new(models.Ipv4)
	/*
		result := C.py_get_ipv4()

		var dwLocalIp uint32 = uint32(result.dwLocalIp)
		var dwNetMask uint32 = uint32(result.dwNetMask)
		var dwGateWay uint32 = uint32(result.dwGateWay)
		var dwFirstDns uint32 = uint32(result.dwFirstDns)
		var dwSecDns uint32 = uint32(result.dwSecDns)

		ipv4.DwLocalIp = libs.Long2ip(dwLocalIp)
		ipv4.DwNetMask = libs.Long2ip(dwNetMask)
		ipv4.DwGateWay = libs.Long2ip(dwGateWay)
		ipv4.DwFirstDns = libs.Long2ip(dwFirstDns)
		ipv4.DwSecDns = libs.Long2ip(dwSecDns)
	*/

	ipv4.DwLocalIp = ""
	ipv4.DwNetMask = ""
	ipv4.DwGateWay = ""
	ipv4.DwFirstDns = ""
	ipv4.DwSecDns = ""

	return ipv4

}
