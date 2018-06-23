package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fields struct {
	Id       string `json:"id"`
	Var      string `json:"var"`
	Fn       string `json:"fn"`
	Param1   string `json:"param1"`
	Param2   string `json:"param2"`
	Operator string `json:"operator"`
	Ds       string `json:"ds"`
	Url      string `json:"url"`
	Out      string `json:"out"`   //internal state
	State    string `json:"state"` //internal state

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
