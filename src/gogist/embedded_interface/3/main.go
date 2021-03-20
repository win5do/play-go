package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Array1 []int

func (arr Array1) Len() int {
	return len(arr)
}

func (arr Array1) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array1) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Array2 []int

func (arr Array2) Len() int {
	return len(arr)
}

func (arr Array2) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array2) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Sortable struct {
	sort.Interface
	// other field
	Type string
}

func NewSortable(i sort.Interface) Sortable {
	t := reflect.TypeOf(i).String()

	return Sortable{
		Interface: i,
		Type:      t,
	}
}

func DoSomething(s Sortable) {
	fmt.Println(s.Type)
	fmt.Println(s.Len())
	fmt.Println(s.Less(0, 1))
}

func main() {
	arr1 := Array1{1, 2, 3}
	arr2 := Array2{3, 2, 1, 0}

	DoSomething(NewSortable(arr1))
	DoSomething(NewSortable(arr2))
}
