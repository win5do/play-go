package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(replaceBlank("we are happy."))
	fmt.Println(replaceBlank(" we are happy.  "))
	fmt.Println(replaceBlank("we_are_happy."))
	fmt.Println(replaceBlank(""))

	fmt.Println(strings.ReplaceAll("we are happy.", " ", "%20"))
}
