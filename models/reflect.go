package models

import (
	"errors"
	"fmt"
	"reflect"
)

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

type Result struct {
	Status string
	Code   string
	Msg    string
}

func Exec(payload interface{}) Result {
	//reflection
	funcs := map[string]interface{}{
		"foo": foo,
		"bar": bar,
		"add": add,
		"sub": sub,
	}

	arr := payload.(map[string]interface{})
	fmt.Println("===>", arr["fn"])
	fmt.Println("===>", arr["param1"])
	fmt.Println("===>", arr["param2"])

	x, _ := Call(funcs, arr["fn"].(string), arr["param1"], arr["param2"])
	y := reflect.ValueOf(x).Interface().([]reflect.Value)
	z := y[0].Interface().(string)

	out := Result{"Success", "1", z}
	return out
}
