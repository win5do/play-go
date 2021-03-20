package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	Sex  string
}

func (s *People) Add(a, b int) int {
	return a + b
}

func (s *People) GetName() string {
	return s.Name
}

func (s *People) SetAge(age int) {
	s.Age = age
}

func reflectStruct(a interface{}) {
	ptr := reflect.ValueOf(a)
	val := reflect.Indirect(ptr)
	typ := val.Type()

	num := val.NumField()
	fmt.Printf("struct have %d field\n", num)

	for i := 0; i < num; i++ {
		// get name and value
		fmt.Printf("filed name: %v value: %v\n", typ.Field(i).Name, val.Field(i))

		// get tag
		tagVal := typ.Field(i).Tag.Get("json")
		fmt.Printf("filed tag=%v\n", tagVal)
	}

	// get method
	numOfMethod := ptr.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	for i := 0; i < numOfMethod; i++ {
		fmt.Printf("sturct method: %v\n", ptr.Type().Method(i).Name)
	}

	// method sort by name (ASCII order)

	// call method 2
	ptr.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(5))
	params = append(params, reflect.ValueOf(10))
	// call method 1
	res := ptr.Method(0).Call(params)
	fmt.Printf("result: %v\n", res[0].Int())

}

func main() {
	var a = &People{
		Name: "foo",
		Age:  33,
		Sex:  "man",
	}
	reflectStruct(a)
}
