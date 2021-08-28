package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(reverseSentence("1234"))
	fmt.Println(reverseSentence("I am a student."))
	fmt.Println(reverseSentence("I am a student. "))
	fmt.Println(reverseSentence("你 好"))
	fmt.Println(reverseSentence(" "))
}
