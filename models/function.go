package models

import (
	"fmt"
	"strconv"
)

func (f Fields) Add(infos map[string]interface{}) (result int, err error) {
	//do something
	fmt.Println("running function:", f.Fn)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)
	a, err := strconv.Atoi(f.Param1)
	b, err := strconv.Atoi(f.Param2)

	x := a + b
	result = x
	infos[f.Var] = result

	return
}
func (f Fields) Sub(infos map[string]interface{}) (result int, err error) {
	//do something
	fmt.Println("running function:", f.Fn)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)
	a, err := strconv.Atoi(f.Param1)
	b, err := strconv.Atoi(f.Param2)

	x := a - b
	//result = strconv.Itoa(x)
	result = x
	infos[f.Out] = result
	infos[f.State] = 1
	return
}
