package main

import (
	"fmt"
	"reflect"
)

type myError struct {
	Err error
}

func (e *myError) Error() string { return e.Err.Error() }

func isNil(i interface{}) bool {
	return reflect.TypeOf(i) == nil && !reflect.ValueOf(i).IsValid()
}

func main() {
	var e *myError
	fmt.Println(isNil(e)) // false

	var i interface{}
	fmt.Println(isNil(i)) // true
}
