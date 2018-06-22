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

	resp["count"] = 0
	resp["desc"] = "-"
	resp["success"] = false
	var xdata []map[string]string
	resp["data"] = xdata

	//final flow
	fmt.Println("flow:", flow)
	resp["data"] = flow

	if id == "" {
		resp["desc"] = "config not found"
		q.Data["json"] = resp
		q.ServeJSON()
		return
	}

	/*str := `{"page": 1, "fruits": ["apple", "peach"]}`
	serv := jsonflow{}
	json.Unmarshal([]byte(str), &serv)
	fmt.Println(serv)
	fmt.Println(serv.Nodes[0])
	*/

	//var dat map[string]interface{}
	var f interface{}
	byt := []byte(flow)
	if err := json.Unmarshal(byt, &f); err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	arr := m["action"].([]interface{})
	//fmt.Println("===>", arr[0])

	for i := 0; i < len(arr); i++ {
		fmt.Println("array:", arr[i])
		function := arr[i].(map[string]interface{})
		fmt.Println("===>", function["fn"])

		out := models.Exec(arr[i])
		fmt.Println("result", out.Msg)

	}
	/*for k, v := range arr {
		fmt.Println("array:", k, v)
		abc := arr[0].(map[string]interface{})
		fmt.Println("===>", abc["fn"])

	}*/
	/*for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	*/

	//var out models.Result
	//out := models.Exec(flow)
	//fmt.Println("result", out.Msg)

	out := "ok"
	q.Data["json"] = out
	beego.Info(out)
	q.ServeJSON()

}

type jsonflow struct {
	Type  string        `json:"page"`
	Nodes []interface{} `json:"fruits"`
}
