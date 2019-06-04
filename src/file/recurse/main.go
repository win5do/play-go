package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanFile(path string, fileMap map[string][]byte) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("read dir err: %v, path: %v", err, path)
		return
	}

	for _, fi := range fs {
		fileName := filepath.Join(path, fi.Name())

		if fi.IsDir() {
			// 目录, 递归遍历
			scanFile(fileName, fileMap)
		} else {
			bts, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Printf("read file err: %v", err)
				continue
			}
			fileMap[fileName] = bts
		}
	}
}

func main() {
	fileMap := make(map[string][]byte)
	scanFile("src", fileMap)
	for k, _ := range fileMap {
		fmt.Printf("%v\n", k)
	}
}
