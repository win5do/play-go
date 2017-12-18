package main

import (
	"os"
	"archive/tar"
	"fmt"
)

func main() {
	htmlTar, _ := os.Open("address_test.html.tar.gz")
	tr := tar.NewReader(htmlTar)
	htmlByte := []byte{}
	tr.Read(htmlByte)
	fmt.Println(htmlByte)
}
