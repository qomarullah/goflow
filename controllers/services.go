package controllers

import (
	"encoding/json"
	"fmt"
	"goflow/lib"
	"io/ioutil"
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

	flowtype := beego.AppConfig.String("flow.type")
	//Type file json
	if flowtype == "files" {

		//flow, err := getPages(id)
		strflow, err := getPagesString(id)
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

			strflow = strings.Replace(strflow, "["+strkeys[i]+"]", mymap[strkeys[i]][0], 5)
			fmt.Println(strkeys[i], "==>", mymap[strkeys[i]][0])
		}
		//end

		//start unmarshal
		flow := new(Page)
		json.Unmarshal([]byte(strflow), &flow)
		i := 0
		//for _, field := range flow.Action {
		task := lib.Task{}
		task.Info = infos //set hashmap container

		valid := true
		for i = 0; i < len(flow.Action) && valid; i++ {

			task.Status = 1
			field := flow.Action[i]
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
				i = len(flow.Action)
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

type Page struct {
	Service string `json:"service"`
	Action  []lib.Fields
}

func (p Page) pageToString() string {
	return lib.ToJson(p)
}

func getPages(file string) (c Page, err error) {
	raw, err := ioutil.ReadFile("files/" + file + ".json")

	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(1)
		return c, err
	}

	json.Unmarshal(raw, &c)
	return c, err
}
func getPagesString(file string) (c string, err error) {
	raw, err := ioutil.ReadFile("files/" + file + ".json")

	if err != nil {
		fmt.Println(err.Error())
		return c, err
	}

	c = string(raw)
	return c, err
}

type Result struct {
	Status     string
	ResultCode int
	ResultDesc string
	Msg        interface{}
}
