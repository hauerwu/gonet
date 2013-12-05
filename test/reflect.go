package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	id int
}

func main() {
	s := make([]Test,10,10)

	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Slice{
		fmt.Println(t.Elem())
	}
}
