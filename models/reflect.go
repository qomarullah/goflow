package models

import (
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
func Exec(task Fields, infos map[string]interface{}) Result {

	out := Result{"Failed", "1", "nothing"}

	x, err := Invoke(task, task.Fn, infos)
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

//func Populate(task Fields) Fields {

//}
