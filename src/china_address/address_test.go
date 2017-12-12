package main

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	bodyByte :=	tryHttp(1)
	fmt.Println(string(bodyByte))
}
