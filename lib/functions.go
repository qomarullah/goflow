package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (f Fields) Set(task Task) Task {
	//do something
	infos := task.Info
	fmt.Println("running function:", f.Fn)
	fmt.Println("got var:", f.Var)
	fmt.Println("got val:", f.Val)

	val := Populate(f.Val, infos)

	fmt.Println("put", f.Var, f.Val)
	infos[f.Var] = val
	task.Info = infos
	fmt.Println("print", task.Info[f.Val])

	return task
}
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
			fmt.Println("EXIT-IF-ERR", task.Step)
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

func (f Fields) Loop(task Task) Task {
	//do something
	infos := task.Info
	fmt.Println("running function:", f.Loop)
	fmt.Println("got start:", f.Start)
	fmt.Println("got end:", f.End)
	fmt.Println("got true:", f.True)
	fmt.Println("got sleep:", f.Delay)

	start1 := Populate(f.Start, infos)
	end1 := Populate(f.End, infos)
	start, err := strconv.Atoi(start1)
	end, err := strconv.Atoi(end1)

	delay := 0
	if f.Delay != "" {
		delay, err = strconv.Atoi(f.Delay)
	}
	CheckErr(err, task)

	//set next
	next := f.True

	//initial

	for loop := start; loop < end; loop++ {

		valid := true
		parent_step := task.Step
		fmt.Println("LOOP-----------------", parent_step, loop, "--------------------")

		for i := 0; i < len(next) && valid; i++ {

			fmt.Println("-----------------", parent_step, i, "--------------------")
			task.Status = 1
			field := next[i]
			fmt.Println(field.ToString())
			task = Exec(i, field, task)
			if delay > 0 {
				fmt.Println("sleep...")
				time.Sleep(time.Duration(delay) * time.Second)
			}

			fmt.Println("RESULT-LOOP", task, task.Status, task.Resp)
			if task.Err != nil {
				fmt.Println("EXIT-LOOP-ERR", task.Step)
				break
			}

			if task.Status == 1 {

			} else if task.Status == 10 {
				valid = false
				i = len(next)
				fmt.Println("BREAK-LOOP")

			} else {
				task.Status = 0
				fmt.Println("EXIT-LOOP", task.Step)
				break
			}

		}
	}
	task.Status = 1
	fmt.Println("FINISH-LOOP")
	fmt.Println("RESULT-LOOP", task, task.Status, task.Resp)

	return task
}

func (f Fields) Go(task Task) Task {
	//do something
	//infos := task.Info
	fmt.Println("running function:", f.Go)
	fmt.Println("delay:", f.Delay)
	delay, err := strconv.Atoi(f.Delay)

	CheckErr(err, task)
	//set next
	next := f.True

	//initial
	valid := true
	parent_step := task.Step

	for i := 0; i < len(next) && valid; i++ {

		fmt.Println("-----------------", parent_step, i, "--------------------")
		task.Status = 1
		field := next[i]
		fmt.Println(field.ToString())
		//task = Exec(i, field, task)
		go Exec(i, field, task)
		fmt.Println("GO...", field, task)

		if delay > 0 {
			fmt.Println("sleep...")
			time.Sleep(time.Duration(delay) * time.Second)
		}

	}
	task.Status = 1
	fmt.Println("FINISH-GO")
	fmt.Println("RESULT-GO", task, task.Status, task.Resp)

	return task
}
