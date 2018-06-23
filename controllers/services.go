package controllers

import (
	"encoding/json"
	"fmt"
	"goflow/models"
	"io/ioutil"
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

type Infos struct {
	KeyVal []KeyVal
}

type KeyVal struct {
	Key string
	Val string
}

func (q *ServicesController) Services() {
	id := q.GetString("id")
	out := models.Result{}

	//var infos map[srting]string
	infos := make(map[string]interface{})
	infos["test"] = "ok"
	fmt.Println("get", infos["test"])

	if id == "" {
		out.Status = "Failed"
		out.ResultCode = "-1"
		out.Msg = "ID Not Found"

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
			out.ResultCode = "-1"
			out.Msg = "ID Not Found"

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

		for _, task := range flow.Action {
			fmt.Println(task.ToString())
			out = models.Exec(task, infos)
			fmt.Println("RESULT", out)

		}
	}

	q.Data["json"] = out
	q.ServeJSON()

}

type Page struct {
	Service string `json:"service"`
	Action  []models.Fields
}

func (p Page) pageToString() string {
	return models.ToJson(p)
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

	//json.Unmarshal(raw, &c)
	c = string(raw)
	return c, err
}
