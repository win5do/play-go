package main

import (
	"regexp"
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
	"encoding/json"
	. "../throw"
	"os"
)

type Address struct {
	Value    string              `json:"value"`
	Label    string              `json:"label"`
	Children map[string]*Address `json:"children, omitempty"`
}

const URL = "http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201703/t20170310_1471429.html"

// 重试get
func tryHttp(times int) []byte {
	var client http.Client
	var resp *http.Response
	err := errors.New("init")
	for i := 0; err != nil && i < times; i++ {
		req, err1 := http.NewRequest("GET", URL, nil)
		err = err1
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
		req.Header.Set("Host", "www.stats.gov.cn")
		req.Close = true
		resp, err = client.Do(req)
	}
	Throw(err, "")
	bodyByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	Throw(err, "")
	ioutil.WriteFile("address.html", bodyByte, 0777)
	return bodyByte
}

// 正则匹配出数据
func transData() map[string]*Address {
	bodyByte := tryHttp(3)
	body := string(bodyByte)
	reg := regexp.MustCompile(`<p class="MsoNormal">.*?<span lang="EN-US">(\d+)<span>.+?<span style="font-family: 宋体">[\p{Zs}\s]*(\p{Han}+)</span>.*?</p>`)
	matched := reg.FindAllStringSubmatch(body, -1)
	addressMap := make(map[string]*Address)
	for _, item := range matched {
		var p, c, a bool
		var err error
		Value := item[1]
		Label := item[2]
		fmt.Println(Value)
		p, err = regexp.MatchString(`^\d{2}0000$`, Value)
		pv := Value[0:2] + "0000"
		if p {
			addressMap[pv] = &Address{Value: Value, Label: Label}
			continue
		}

		c, err = regexp.MatchString(`^\d{2}([0-9][1-9]|[1-9][0-9])00$`, Value)
		cv := Value[0:4] + "00"
		if c {
			if len(addressMap[pv].Children) == 0 {
				addressMap[pv].Children = make(map[string]*Address)
			}
			addressMap[pv].Children[cv] = &Address{Value: Value, Label: Label}
			continue
		}
		a, err = regexp.MatchString(`^\d{4}00$`, Value)
		if !a {
			if len(addressMap[pv].Children[cv].Children) == 0 {
				addressMap[pv].Children[cv].Children = make(map[string]*Address)
			}
			addressMap[pv].Children[cv].Children[Value] = &Address{Value: Value, Label: Label}
			continue
		}
		Throw(err, Value)
	}
	return addressMap
}

func main() {
	mapData := transData()
	fmt.Println(mapData)
	file, _ := os.Create("output.json")
	enc := json.NewEncoder(file)
	err := enc.Encode(mapData)
	Throw(err, "")
}
