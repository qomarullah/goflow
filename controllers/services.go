package controllers

import (
	"encoding/json"
	"fmt"
	"goflow/models"

	"reflect"
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

	//flow := `{"type":"flow","action":[{"fn":"add","param1":"[param1]","param2":"[param2]"},{"fn":"sub","param1":"[param1]","param2":"[param2]"}]}`
	flow := beego.AppConfig.String("flow." + id)
	fmt.Println("=======>" + flow)
	mymap := q.Ctx.Request.URL.Query()
	keys := reflect.ValueOf(mymap).MapKeys()
	strkeys := make([]string, len(keys))

	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
		flow = strings.Replace(flow, "["+strkeys[i]+"]", mymap[strkeys[i]][0], 5)
		fmt.Println(strkeys[i], "==>", mymap[strkeys[i]][0])
	}

	var resp map[string]interface{}
	if resp == nil {
		resp = make(map[string]interface{})
	}

	//final flow
	fmt.Println("flow:", flow)
	resp["data"] = flow

	if id == "" {
		resp["desc"] = "config not found"
		q.Data["json"] = resp
		q.ServeJSON()
		return
	}

	var f interface{}
	byt := []byte(flow)
	if err := json.Unmarshal(byt, &f); err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	arr := m["action"].([]interface{})

	//default
	out := models.Result{}

	for i := 0; i < len(arr); i++ {
		fmt.Println("array:", arr[i])
		jsonString, _ := json.Marshal(arr[i])
		out = models.Exec(jsonString)
		fmt.Println("RESULT", out)
	}

	//out = "ok"
	q.Data["json"] = out
	beego.Info(out)
	q.ServeJSON()

}
