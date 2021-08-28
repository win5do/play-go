package main

import (
	"fmt"
	"reflect"
)

func toSlice(itf interface{}) []interface{} {
	docs := reflect.ValueOf(itf)
	if docs.Kind() != reflect.Slice {
		return nil
	}

	l := docs.Len()
	temp := make([]interface{}, l)
	for i := 0; i < l; i++ {
		temp[i] = docs.Index(i).Interface()
	}

	return temp
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(toSlice(arr))
}
