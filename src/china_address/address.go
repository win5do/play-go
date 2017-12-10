package main

import (
	//"encoding/json"
	//"os"
	//"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
	. "../throw"
	"fmt"
)

type address struct {
	value    string
	label    string
	children []address
}

func getOriginData() {
	res, err := http.Get("http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201703/t20170310_1471429.html")
	Throw(err, "")
	defer res.Body.Close()
	bodyByte, err := ioutil.ReadAll(res.Body)
	Throw(err, "")
	body := string(bodyByte)
	reg := regexp.MustCompile(`<p class="MsoNormal">.+?<span lang="EN-US">(\d+)<span>.+?<span style="font-family: 宋体">[\p{Zs}\s]*(\p{Han}+)</span>.*?</p>`)
	matched := reg.FindAllStringSubmatch(body, -1)
	//fmt.Println(matched)
	var result []address
	for _, item := range matched {
		result = append(result, address{value: item[1], label: item[2]})
	}
	fmt.Println(result)
}

func main() {
	getOriginData()
}

//func main() {
//	v := []int{1, 2, 3, 4, 5}
//	file, _ := os.Create("output.json")
//	fmt.Println(file)
//	encoder := json.NewEncoder(file)
//	err := encoder.Encode(v)
//	if err != nil {
//		panic(err)
//	}
//}
