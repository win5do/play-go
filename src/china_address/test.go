package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	var err error
	file, _ := os.Open("test.html.gz")
	defer file.Close()
	zr, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	//defer zr.Close()

	//b := new(bytes.Buffer)
	b := bytes.Buffer{}
	b.ReadFrom(zr)

	fmt.Println(b.String())
}
