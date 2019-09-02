package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	pHead := deserialize("1,2,4,$,$,$,3,5,$,$,6,$,$")
	str := serialize(pHead)
	fmt.Println(str)
}
