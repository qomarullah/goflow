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
	CheckErr(err, task)

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

func (f Fields) If(task Task) Task {
	//do something
	infos := task.Info
	fmt.Println("running function:", f.Fn)
	fmt.Println("got opr:", f.Opr)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)
	fmt.Println("got negation:", f.Neg)

	//initial
	cond := false
	if strings.EqualFold(f.Opr, "<") {
		a1 := Populate(f.Param1, infos)
		b1 := Populate(f.Param2, infos)
		a, err := strconv.Atoi(a1)
		b, err := strconv.Atoi(b1)

		CheckErr(err, task)

		if a < b {
			cond = true
		}

	}
	if strings.EqualFold(f.Opr, "<=") {
		a1 := Populate(f.Param1, infos)
		b1 := Populate(f.Param2, infos)
		a, err := strconv.Atoi(a1)
		b, err := strconv.Atoi(b1)

		CheckErr(err, task)

		if a <= b {
			cond = true
		}

	}
	if strings.EqualFold(f.Opr, ">") {
		a1 := Populate(f.Param1, infos)
		b1 := Populate(f.Param2, infos)
		a, err := strconv.Atoi(a1)
		b, err := strconv.Atoi(b1)

		CheckErr(err, task)

		if a > b {
			cond = true
		}

	}
	if strings.EqualFold(f.Opr, ">=") {
		a1 := Populate(f.Param1, infos)
		b1 := Populate(f.Param2, infos)
		a, err := strconv.Atoi(a1)
		b, err := strconv.Atoi(b1)

		CheckErr(err, task)

		if a >= b {
			cond = true
		}

	}
	if strings.EqualFold(f.Opr, "eq") {
		a := Populate(f.Param1, infos)
		b := Populate(f.Param2, infos)

		if strings.EqualFold(a, b) {
			cond = true
		}

	}
	// negation
	if f.Neg == "true" {
		cond = !cond
	}
	//set next
	next := f.True
	if cond == false {
		next = f.False
	}
	valid := true
	parent_step := task.Step
	for i := 0; i < len(next) && valid; i++ {
		fmt.Println("-----------------", parent_step, i, "--------------------")
		task.Status = 1
		field := next[i]
		fmt.Println(field.ToString())
		task = Exec(i, field, task)
		fmt.Println("RESULT-IF", task, task.Status, task.Resp)
		if task.Err != nil {
			fmt.Println("EXIT-IF", task.Step)
			break
		}

		if task.Status == 1 {

		} else if task.Status == 10 {
			valid = false
			i = len(next)
			fmt.Println("BREAK-IF")

		} else {
			task.Status = 0
			fmt.Println("EXIT-IF", task.Step)
			break
		}

	}
	fmt.Println("FINISH-IF")
	fmt.Println("RESULT-IF", task, task.Status, task.Resp)

	return task
}

func CheckErr(err error, task Task) Task {
	if err != nil {
		task.Err = err
		task.Status = -3
		return task
	}
	return task
}
