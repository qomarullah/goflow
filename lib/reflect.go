package lib

import (
	"errors"
	"fmt"
	"reflect"
)

type Task struct {
	Step   int
	Status int
	Err    error
	Info   map[string]string
	Resp   interface{}
}

/*func Invoke(m interface{}, name string, params ...interface{}) (resp []reflect.Value, err error) {
	fmt.Println("masuk.....")
	f := reflect.ValueOf(m)
	fmt.Println("function", name)
	fmt.Println("valuexxx", params[0], len(params))

	if f.MethodByName(name).IsValid() == false {
		err = errors.New("Function not found")
		return
	}
	inputs := make([]reflect.Value, len(params))
	for i, _ := range params {
		inputs[i] = reflect.ValueOf(params[i])
	}

	resp = f.MethodByName(name).Call(inputs)
	fmt.Println("invoke", resp[0], resp[1])

	return

}
*/
func Exec(step int, field Fields, task Task) Task {

	task.Step = step
	fn := field.Fn
	f := reflect.ValueOf(field)
	fmt.Println("function", fn)
	fmt.Println("value", task, 1)

	if f.MethodByName(fn).IsValid() == false {

		err := errors.New("Function not found")
		task.Status = -2
		task.Err = err
		return task
	}
	inputs := make([]reflect.Value, 1)
	/*for i, _ := range params {
		inputs[i] = reflect.ValueOf(task)
	}
	*/
	inputs[0] = reflect.ValueOf(task)
	fmt.Println("invoke-start", task)

	resp := f.MethodByName(fn).Call(inputs)
	task = resp[0].Interface().(Task)
	//task = y
	//y := reflect.ValueOf(resp[0]).Interface().(Task)
	fmt.Println("invoke-end", task, task.Status)

	/*y := reflect.ValueOf(resp).Interface().([]reflect.Value)
	z := y[0].Interface().(Task)
	fmt.Println("invoke", z)

	if z.Err != nil {
		z.Status = -1

	}
	task = z
	*/
	return task
}

//func Populate(task Fields) Fields {

//}
