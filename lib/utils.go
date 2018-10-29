package lib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Fields struct {
	Id      string   `json:"id"`
	Var     string   `json:"var"`
	Val     string   `json:"val"`
	Fn      string   `json:"fn"`
	Param1  string   `json:"param1"`
	Param2  string   `json:"param2"`
	Neg     string   `json:"neg"`
	Opr     string   `json:"opr"`
	Ds      string   `json:"ds"`
	Url     string   `json:"url"`
	Out     string   `json:"out"`   //internal state
	State   string   `json:"state"` //internal state
	True    []Fields `json:"true"`
	False   []Fields `json:"false"`
	Start   string   `json:"start"`
	End     string   `json:"end"`
	Delay   string   `json:"delay"`
	Type    string   `xml:"type,attr"`
	Timeout string   `xml:"timeout,attr"`
	Body    string   `xml:"body,attr"`
}

type FieldsXML struct {
	Id     string `xml:"id,attr"`
	Var    string `xml:"var,attr"`
	Val    string `xml:"val,attr"`
	Fn     string `xml:"fn,attr"`
	Param1 string `xml:"param1,attr"`
	Param2 string `xml:"param2,attr"`
	Neg    string `xml:"neg,attr"`
	Opr    string `xml:"opr,attr"`
	Ds     string `xml:"ds,attr"`

	Url   string `xml:"url,attr"`
	Out   string `xml:"out,attr"`   //internal state
	State string `xml:"state,attr"` //internal state
	//True   []FieldsXML `xml:"true"`
	//False  []FieldsXML `xml:"false"`
	Start   string      `xml:"start,attr"`
	End     string      `xml:"end,attr"`
	Delay   string      `xml:"delay,attr"`
	Next    []FieldsXML `xml:"go"`
	Type    string      `xml:"type,attr"`
	Timeout string      `xml:"timeout,attr"`
	Body    string      `xml:"body,attr"`
}

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}

	return string(bytes)
}

func (f Fields) ToString() string {
	return ToJson(f)
}
func (f FieldsXML) ToString() string {
	return ToJson(f)
}
func Populate(x string, infos map[string]string) (y string) {

	y = x
	if strings.HasPrefix(x, "[") && strings.HasSuffix(x, "]") {
		xx := strings.TrimPrefix(x, "[")
		xx = strings.TrimSuffix(xx, "]")
		y = infos[xx]
		fmt.Println("populate-res", "x:"+xx, "y:"+y)
		//looping
		if strings.HasPrefix(x, "[") && strings.HasSuffix(x, "]") {
			Populate(y, infos)
		}
	}
	return
}

func CheckErr(err error, task Task) Task {
	if err != nil {
		task.Err = err
		task.Status = -3
		return task
	}
	return task
}

type Page struct {
	Service string   `json:"service"`
	GO      []Fields `json:"go"`
}

type PageXML struct {
	Service xml.Name    `xml:"service"`
	GO      []FieldsXML `xml:"go"`
}
type PageXML2 struct {
	GO []FieldsXML `xml:"go"`
}

func (p Page) pageToString() string {
	return ToJson(p)
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

func GetPagesString(file, format string) (c string, err error) {
	raw, err := ioutil.ReadFile("files/" + file + "." + format)

	if err != nil {
		fmt.Println(err.Error())
		return c, err
	}

	c = string(raw)
	return c, err
}
