package models

import (
	"fmt"
	"strconv"
)

func foo() string {
	fmt.Println("we are running foo")
	return "0"
}

func bar(a, b, c string) string {
	fmt.Println("we are running bar", a, b, c)
	return "10"
}

func add(a, b string) string {
	a1, _ := strconv.Atoi(a)
	b1, _ := strconv.Atoi(b)
	x := a1 + b1
	out := strconv.Itoa(x)
	return out
}
func sub(a, b string) string {
	a1, _ := strconv.Atoi(a)
	b1, _ := strconv.Atoi(b)
	x := a1 - b1
	out := strconv.Itoa(x)
	return out
}
