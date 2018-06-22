package models

import (
	"fmt"
	"strconv"
)

type Fields struct {
	Fn       string `json:"fn"`
	Param1   string `json:"param1"`
	Param2   string `json:"param2"`
	Operator string `json:"operator"`
	Ds       string `json:"ds"`
	Url      string `json:"url"`
}

func (f Fields) Add() (result int, err error) {
	//do something
	fmt.Println("running function:", f.Fn)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)
	a, err := strconv.Atoi(f.Param1)
	b, err := strconv.Atoi(f.Param2)

	x := a + b
	//result = strconv.Itoa(x)
	result = x
	return
}
func (f Fields) Sub() (result int, err error) {
	//do something
	fmt.Println("running function:", f.Fn)
	fmt.Println("got param1:", f.Param1)
	fmt.Println("got param2:", f.Param2)
	a, err := strconv.Atoi(f.Param1)
	b, err := strconv.Atoi(f.Param2)

	x := a - b
	//result = strconv.Itoa(x)
	result = x
	return
}
