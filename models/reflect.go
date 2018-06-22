package models

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Result struct {
	Status     string
	ResultCode string
	Msg        interface{}
}

func Invoke(m interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {

	f := reflect.ValueOf(m)

	inputs := make([]reflect.Value, len(params))
	/*for i, _ := range params {
		inputs[i] = reflect.ValueOf(params[i])
	}
	*/
	result = f.MethodByName(name).Call(inputs)
	fmt.Println(result)
	return

}
func Exec(task []byte) Result {

	req := Fields{}
	json.Unmarshal(task, &req)
	//fmt.Println(req)
	//fmt.Println(req.Fn)

	out := Result{"Failed", "1", "ok"}
	x, err := Invoke(req, req.Fn)
	if err != nil {
		out.Msg = err.Error()
		fmt.Println("failed", out.Msg)
		return out
	}

	y := reflect.ValueOf(x).Interface().([]reflect.Value)
	z := y[0].Interface()

	out = Result{"Success", "1", z}
	return out
}
