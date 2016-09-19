package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
	"webTest/jobs"
	"webTest/libs"
	"webTest/models"
)

type TaskController struct {
	BaseController
}

// 添加任务
func (this *TaskController) Add() {

	if this.isPost() {

		task := new(models.Task)
		task.Id = 1

		task_type, _ := this.GetInt("task_name")

		if task_type == 0 { //网口抓包
			command := "tcpdump"
			ip := strings.TrimSpace(this.GetString("ip"))

			if ip != "" {
				command += " host " + ip
			}

			port := strings.TrimSpace(this.GetString("port"))

			if port != "" {
				command += " port " + port
			}

			task.Command = command

		} else {
			task.Command = strings.TrimSpace(this.GetString("command"))
		}

		task.Timeout, _ = this.GetInt("timeout")
		task.Concurrent = 1
		task.Status = 1

		//if task.TaskName == "" || task.Command == "" {
		//	this.ajaxMsg("请填写完整信息", MSG_ERR)
		//}

		//if _, err := models.TaskAdd(task); err != nil {
		//	this.ajaxMsg(err.Error(), MSG_ERR)
		//}

		job, err := jobs.NewJobFromTask(task)

		if err != nil {
			this.showMsg(err.Error())
		}

		job.Run()
		this.ajaxMsg("", MSG_OK)
		//this.redirect(beego.URLFor("TaskController.ViewLog"))

	}

	// 分组列表

	this.Data["pageTitle"] = "添加任务"
	this.display()
}

// 任务执行日志列表
func (this *TaskController) Logs() {
	//taskId, _ := this.GetInt("id")
	taskId := 1
	page, _ := this.GetInt("page")

	if page < 1 {
		page = 1
	}

	/**task, err := models.TaskGetById(taskId)
	if err != nil {
		this.showMsg(err.Error())
	}**/

	result, count := models.TaskLogGetList(page, this.pageSize, "task_id", 1)

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["start_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["process_time"] = float64(v.ProcessTime) / 1000
		row["ouput_size"] = libs.SizeFormat(float64(len(v.Output)))
		row["status"] = v.Status
		list[k] = row
	}

	this.Data["pageTitle"] = "任务执行日志"
	this.Data["list"] = list
	//this.Data["task"] = task
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("TaskController.Logs", "id", taskId), true).ToString()
	this.display()
}

// 查看日志详情
func (this *TaskController) ViewLog() {
	//id, _ := this.GetInt("id")
	id := models.TaskLogListId()
	taskLog, err := models.TaskLogGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	//task, err := models.TaskGetById(taskLog.TaskId)
	//if err != nil {
	//	this.showMsg(err.Error())
	//}
	println(taskLog.Output)
	data := make(map[string]interface{})
	data["id"] = taskLog.Id
	data["output"] = taskLog.Output
	data["error"] = taskLog.Error
	data["start_time"] = beego.Date(time.Unix(taskLog.CreateTime, 0), "Y-m-d H:i:s")
	data["process_time"] = float64(taskLog.ProcessTime) / 1000
	data["ouput_size"] = libs.SizeFormat(float64(len(taskLog.Output)))
	data["status"] = taskLog.Status

	//this.Data["task"] = task

	this.Data["data"] = data
	this.Data["pageTitle"] = "查看日志"

	this.display()

}

// 批量操作日志
func (this *TaskController) LogBatch() {
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
			models.TaskLogDelById(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}

// 立即执行
func (this *TaskController) Run() {

	id, _ := this.GetInt("id")

	task, err := models.TaskGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	job, err := jobs.NewJobFromTask(task)
	if err != nil {
		this.showMsg(err.Error())
	}

	job.Run()

	this.redirect(beego.URLFor("TaskController.ViewLog", "id", job.GetLogId()))
}
