package actions

import (
	"fmt"
	"strconv"
	"strings"
)

func (f Fields) Arith(task Task) Task {
	//do something
	infos := task.Info
	fmt.Println("running function:", f.Fn)
	fmt.Println("got opr:", f.Opr)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)

	a1 := Populate(f.Param1, infos)
	b1 := Populate(f.Param2, infos)
	a, err := strconv.Atoi(a1)
	b, err := strconv.Atoi(b1)
	if err != nil {
		task.Err = err
		task.Status = -3
		return task
	}

	var result int
	if strings.EqualFold(f.Opr, "*") {
		result = a * b
	}
	if strings.EqualFold(f.Opr, "/") {
		result = a / b
	}
	if strings.EqualFold(f.Opr, "+") {
		result = a + b
	}
	if strings.EqualFold(f.Opr, "-") {
		result = a - b
	}
	if strings.EqualFold(f.Opr, "<") {
		if a < b {
			result = 1
		}
	}
	if strings.EqualFold(f.Opr, "<=") {
		if a <= b {
			result = 1
		}
	}
	if strings.EqualFold(f.Opr, ">") {
		if a > b {
			result = 1
		}
	}
	if strings.EqualFold(f.Opr, ">=") {
		if a > b {
			result = 1
		}
	}

	res := strconv.Itoa(result)
	fmt.Println("put", f.Out, res)
	infos[f.Out] = res
	task.Info = infos
	fmt.Println("print", task.Info[f.Out])

	return task
}
func (f Fields) Break(task Task) Task {
	//do something
	task.Status = 10
	return task
}
