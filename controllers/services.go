package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"goflow/lib"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about service
type ServicesController struct {
	beego.Controller
}

// @Title Query Services
// @Description service flow from config
// @Param	id		query 	string	true	"Service flow id"
// @Success 200 {string} success
// @Failure 403 data not found
// @router /services [get]

func (q *ServicesController) Services() {
	id := q.GetString("id")
	out := Result{}
	out.Status = "Failed"
	out.ResultCode = -1
	out.ResultDesc = "Not Running"

	infos := make(map[string]string)
	infos["test"] = "ok"
	fmt.Println("get", infos["test"])

	if id == "" {
		out.Status = "Failed"
		out.ResultCode = -1
		out.ResultDesc = "Invalid Parameter"

		q.Data["json"] = out
		q.ServeJSON()
		return
	}

	//Type format
	flow_format := beego.AppConfig.String("flow.format")

	strflow, err := lib.GetPagesString(id, flow_format)
	if err != nil {
		out.Status = "Failed"
		out.ResultCode = -1
		out.ResultDesc = "ID Not Found"

		q.Data["json"] = out
		q.ServeJSON()
		return
	}

	//replace predefine params
	mymap := q.Ctx.Request.URL.Query()
	keys := reflect.ValueOf(mymap).MapKeys()
	strkeys := make([]string, len(keys))

	fmt.Println(strflow)
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()

		strflow = strings.Replace(strflow, "["+strkeys[i]+"]", mymap[strkeys[i]][0], -1)
		fmt.Println(strkeys[i], "==>", mymap[strkeys[i]][0])
	}
	//end

	//start unmarshal json/xml

	//var flow interface{}
	if flow_format == "json" {
		flow := &lib.Page{}

		err = json.Unmarshal([]byte(strflow), &flow)
		if err != nil {
			out.Status = "Failed"
			out.ResultCode = -1
			out.ResultDesc = "Failed to parse"

			q.Data["json"] = out
			q.ServeJSON()
			return
		}

		i := 0
		task := lib.Task{}
		task.Info = infos //set hashmap container

		valid := true
		for i = 0; i < len(flow.GO) && valid; i++ {

			task.Status = 1
			field := flow.GO[i]
			fmt.Println(field.ToString())
			task.Step = i
			fmt.Println("-----------------", i, "--------------------")
			task = lib.Exec(i, field, task)
			fmt.Println("RESULT", task, task.Status, task.Resp)
			if task.Err != nil {
				fmt.Println("EXIT", task.Step)
				out.ResultCode = task.Status
				out.ResultDesc = "task-" + strconv.Itoa(i) + "=" + task.Err.Error()
				break
			}

			if task.Status == 1 {

			} else if task.Status == 10 {
				out.ResultDesc = "task-" + strconv.Itoa(i)
				i = len(flow.GO)
				out.ResultCode = task.Status
				out.ResultDesc = "Break:" + strconv.Itoa(i)
				out.Msg = task.Info["resp"]
				fmt.Println("BREAK")

			} else {
				fmt.Println("EXIT", task.Step)
				out.ResultCode = task.Status
				out.ResultDesc = "task-" + strconv.Itoa(i) + "=" + task.Err.Error()
				break
			}

		}
		fmt.Println("FINISH")
		fmt.Println("RESULT", task, task.Status, task.Resp)
		out.Status = "Success"
		out.ResultCode = task.Status
		out.ResultDesc = "Finish"
		out.Msg = task.Info["resp"]

	} else {
		fmt.Println("==>", "xml", strflow)
		//flow := &PageXML{}
		flow := new(lib.PageXML)

		err = xml.Unmarshal([]byte(strflow), &flow)
		if err != nil {
			out.Status = "Failed"
			out.ResultCode = -1
			out.ResultDesc = err.Error()

			q.Data["json"] = out
			q.ServeJSON()
			return
		}

		i := 0
		task := lib.Task{}
		task.Info = infos //set hashmap container

		valid := true
		for i = 0; i < len(flow.GO) && valid; i++ {

			task.Status = 1
			field := flow.GO[i]
			fmt.Println(field.ToString())
			task.Step = i
			fmt.Println("-----------------", i, "--------------------")
			task = lib.ExecXML(i, field, task)
			fmt.Println("RESULT", task, task.Status, task.Resp)
			if task.Err != nil {
				fmt.Println("EXIT", task.Step)
				out.ResultCode = task.Status
				out.ResultDesc = "task-" + strconv.Itoa(i) + "=" + task.Err.Error()
				break
			}

			if task.Status == 1 {

			} else if task.Status == 10 {
				out.ResultDesc = "task-" + strconv.Itoa(i)
				i = len(flow.GO)
				out.ResultCode = task.Status
				out.ResultDesc = "Break:" + strconv.Itoa(i)
				out.Msg = task.Info["resp"]
				fmt.Println("BREAK")

			} else {
				fmt.Println("EXIT", task.Step)
				out.ResultCode = task.Status
				out.ResultDesc = "task-" + strconv.Itoa(i) + "=" + task.Err.Error()
				break
			}

		}

		fmt.Println("FINISH")
		fmt.Println("RESULT", task, task.Status, task.Resp)
		out.Status = "Success"
		out.ResultCode = task.Status
		out.ResultDesc = "Finish"
		out.Msg = task.Info["resp"]

	}

	q.Data["json"] = out
	q.ServeJSON()

}

type Result struct {
	Status     string
	ResultCode int
	ResultDesc string
	Msg        interface{}
}
