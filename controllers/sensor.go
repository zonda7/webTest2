package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
	"time"
	"webTest/libs"
	"webTest/models"
)

type SensorController struct {
	BaseController
}

// 任务列表
func (this *SensorController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	groupId, _ := this.GetInt("groupid")
	filters := make([]interface{}, 0)
	if groupId > 0 {
		filters = append(filters, "group_number", groupId)
	}
	list, count := models.SensorGetList(page, this.pageSize, filters...)

	// 分组列表
	groups, _ := models.TaskGroupGetList(1, 100)

	this.Data["pageTitle"] = "任务列表"
	this.Data["list"] = list
	this.Data["groups"] = groups
	this.Data["groupid"] = groupId
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SensorController.List", "groupid", groupId), true).ToString()
	this.display()
}

// 添加任务
func (this *SensorController) Add() {

	if this.isPost() {
		sensor := new(models.Sensor)
		sensor.SensorName = strings.TrimSpace(this.GetString("sensor_name"))
		sensor.ChannelName = strings.TrimSpace(this.GetString("channel_name"))
		sensor.DecimalPlace, _ = this.GetInt("decimal_place")
		sensor.GroupNumber, _ = this.GetInt("group_id")
		sensor.LimitField = strings.TrimSpace(this.GetString("limit_field"))
		sensor.ItemData = strings.TrimSpace(this.GetString("item_data"))
		sensor.TargetField = strings.TrimSpace(this.GetString("target_field"))
		sensor.NameTable = strings.TrimSpace(this.GetString("name_table"))

		if sensor.SensorName == "" || sensor.ChannelName == "" {
			this.ajaxMsg("请填写完整信息", MSG_ERR)
		}

		_, err := models.SensorAdd(sensor)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}

		this.ajaxMsg("", MSG_OK)

	}

	// 分组列表
	groups, _ := models.TaskGroupGetList(1, 100)
	this.Data["groups"] = groups
	this.Data["pageTitle"] = "添加传感器"
	this.display()
}

//编辑
func (this *SensorController) Edit() {
	id, _ := this.GetInt("id")
	// 根据ID获取传感器
	sensor, err := models.SersorGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {

		println("ss")

		sensor.SensorName = strings.TrimSpace(this.GetString("sensor_name"))
		sensor.ChannelName = strings.TrimSpace(this.GetString("channel_name"))
		sensor.DecimalPlace, _ = this.GetInt("decimal_place")
		sensor.GroupNumber, _ = this.GetInt("group_id")
		sensor.LimitField = strings.TrimSpace(this.GetString("limit_field"))
		sensor.ItemData = strings.TrimSpace(this.GetString("item_data"))
		sensor.TargetField = strings.TrimSpace(this.GetString("target_field"))
		sensor.NameTable = strings.TrimSpace(this.GetString("name_table"))

		if sensor.SensorName == "" || sensor.ChannelName == "" {
			this.ajaxMsg("请填写完整信息", MSG_ERR)
		}

		err := sensor.Update()

		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}

		this.ajaxMsg("", MSG_OK)
	}

	groups, _ := models.TaskGroupGetList(1, 100)

	this.Data["groups"] = groups
	this.Data["aa"] = sensor.GroupNumber

	this.Data["sensor"] = sensor
	this.Data["pageTitle"] = "编辑传器"
	this.display()
}

// 导出传感器
func (this *SensorController) Export() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	groupId, _ := this.GetInt("groupid")
	filters := make([]interface{}, 0)
	if groupId > 0 {
		filters = append(filters, "group_number", groupId)
	}
	list, _ := models.SensorGetList(page, this.pageSize, filters...)

	//w.Write([]string{"编号", "传感器名称", "通道名称", "小数位", "组号", "限定字段", "限定字段值", "目标字段", "数据库表"})

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "传感器名称"
	cell = row.AddCell()
	cell.Value = "通道名称"
	cell = row.AddCell()
	cell.Value = "小数位"
	cell = row.AddCell()
	cell.Value = "组号"
	cell = row.AddCell()
	cell.Value = "限定字段"
	cell = row.AddCell()
	cell.Value = "限定字段值"
	cell = row.AddCell()
	cell.Value = "目标字段"
	cell = row.AddCell()
	cell.Value = "数据库表"

	for _, sensor := range list {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = sensor.SensorName
		cell = row.AddCell()
		cell.Value = sensor.ChannelName
		cell = row.AddCell()
		cell.Value = strconv.Itoa(sensor.DecimalPlace)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(sensor.GroupNumber)
		cell = row.AddCell()
		cell.Value = sensor.LimitField
		cell = row.AddCell()
		cell.Value = sensor.ItemData
		cell = row.AddCell()
		cell.Value = sensor.TargetField
		cell = row.AddCell()
		cell.Value = sensor.NameTable
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	// 分组列表
	//groups, _ := models.TaskGroupGetList(1, 100)

	//this.Data["pageTitle"] = "任务列表"
	//this.Data["list"] = list
	///this.Data["groups"] = groups
	//this.Data["groupid"] = groupId
	//this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SensorController.List", "groupid", groupId), true).ToString()
	//this.display()
}

func (this *SensorController) Lead() {
	excelFileName := "/Users/apple/Documents/go/src/WebTest/MyXLSXFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {

	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			i := 1
			for _, cell := range row.Cells {
				sensor := new(models.Sensor)
				switch i {
				case 1:
					sensor.SensorName = strings.TrimSpace(cell.Value)
				case 2:
					sensor.ChannelName = strings.TrimSpace(cell.Value)
				case 3:
					sensor.DecimalPlace, _ = strconv.Atoi(cell.Value)
				case 4:
					sensor.GroupNumber, _ = strconv.Atoi(cell.Value)
				case 5:
					sensor.LimitField = strings.TrimSpace(cell.Value)
				case 6:
					sensor.ItemData = strings.TrimSpace(cell.Value)
				case 7:
					sensor.TargetField = strings.TrimSpace(cell.Value)
				case 8:
					sensor.NameTable = strings.TrimSpace(cell.Value)

				default:
				}
				_, err := models.SensorAdd(sensor)
				if err != nil {
				}

				i++

			}
		}
	}

}

//上传
func (this *SensorController) Upload() {

	// 获取上传文件
	f, h, err := this.GetFile("editormd-image-file")
	if err == nil {
		// 关闭文件
		f.Close()
	}
	if err != nil {
		// 获取错误则输出错误信息

		this.ajaxMsg(err.Error(), MSG_ERR)
		//this.ServeJson()
		return
	}
	// 获取当前年月
	datePath := time.Now().Format("2006/01")
	// 设置保存目录
	dirPath := "./upload/" + datePath
	// 设置保存文件名
	FileName := h.Filename
	// 将文件保存到服务器中
	err = this.SaveToFile("editormd-image-file", fmt.Sprintf("%s/%s", dirPath, FileName))
	if err != nil {
		// 出错则输出错误信息
		this.ajaxMsg(err.Error(), MSG_ERR)
		//this.ServeJson()
		return
	}

}

func (this *SensorController) Batch() {
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
			models.SersorDel(id)
		}
	}

	this.ajaxMsg("", MSG_OK)

}
