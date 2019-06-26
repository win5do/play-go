package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func funcName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fmt.Println("fullName: ", fullName)
	dotIndex := strings.LastIndex(fullName, ".")
	return fullName[dotIndex+1:]
}

func foo()  {

}

func bar()  {

}

func main()  {
	fmt.Println(funcName(foo))
	fmt.Println(funcName(bar))
}
