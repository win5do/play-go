package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"flag"
	"fmt"
)

var (
	prefix string = "img"
	start  int    = 1
	inDir  string = "in"
	outDir string = "out"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func throw(err error, msg string) {
	if msg != "" {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(msg)
				os.Exit(-1)
			}
		}()
	}

	if err != nil {
		panic(err)
	}
}

func rename() {
	files, err := ioutil.ReadDir("./" + inDir)
	throw(err, inDir+"文件夹不存在")

	if isExist(outDir) {
		os.RemoveAll(outDir)
	}
	os.Mkdir(outDir, 0777)

	for _, file := range files {
		if file.IsDir() || file.Name()[0:1] == "." {
			continue
		}
		fileName := file.Name()
		fileExt := fileName[strings.LastIndex(fileName, "."):]
		fileByte, err := ioutil.ReadFile(inDir + "/" + fileName)
		throw(err, "读取文件出错")
		newFileName := outDir + "/" + prefix + strconv.Itoa(start) + fileExt
		ioutil.WriteFile(newFileName, []byte(fileByte), 0777)
		start++
	}
}

func init() {
	if len(os.Args) > 1 {
		flag.StringVar(&prefix, "prefix", prefix, "重命名后的前缀")
		flag.IntVar(&start, "start", start, "开始的序号，如1")
		flag.StringVar(&inDir, "in", inDir, "需要处理的目录，可以是相对目录或绝对目录，默认当前目录下in")
		flag.StringVar(&outDir, "out", outDir, "输出目录，相对于当前目录，默认当前目录下out")
		flag.Parse()
	} else {
		fmt.Print("请输入重命名前缀:")
		fmt.Scanln(&prefix)
		fmt.Print("请输入起始编号:")
		fmt.Scanln(&start)
		fmt.Print("请输入处理目录:")
		fmt.Scanln(&inDir)
		fmt.Print("请输入输出目录:")
		fmt.Scanln(&outDir)
	}
}

func main() {
	rename()
}
