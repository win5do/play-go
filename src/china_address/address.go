package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"

	. "../throw"
	"compress/gzip"
)

type Address struct {
	Value    string              `json:"value"`
	Label    string              `json:"label"`
	Children map[string]*Address `json:"children,omitempty"`
}

type AddressArray struct {
	Value    string          `json:"value"`
	Label    string          `json:"label"`
	Children []*AddressArray `json:"children,omitempty"`
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
	Throw(err, "请求页面失败")
	bodyByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	Throw(err, "")
	htmlFile, err := os.Create("output/address.html.gz")
	Throw(err, "")
	defer htmlFile.Close()
	zw := gzip.NewWriter(htmlFile)
	_, err = zw.Write(bodyByte)
	Throw(err, "")
	err = zw.Close()
	Throw(err, "")
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

func outputMap(data map[string]*Address) {
	file, err := os.Create("output/map.json")
	Throw(err, "")
	enc := json.NewEncoder(file)
	err = enc.Encode(data)
	Throw(err, "")
}

func mapToArray(data map[string]*Address) []*AddressArray {
	var resArr []*AddressArray
	var sortKey []string
	for k, _ := range data {
		sortKey = append(sortKey, k)
	}
	sort.Strings(sortKey)
	for _, k := range sortKey {
		v := data[k]
		child := mapToArray(v.Children)
		item := AddressArray{
			Value:    v.Value,
			Label:    v.Label,
			Children: child,
		}
		resArr = append(resArr, &item)
	}
	return resArr
}

func outputArray(data map[string]*Address) {
	array := mapToArray(data)
	file, err := os.Create("output/array.json")
	Throw(err, "")
	enc := json.NewEncoder(file)
	err = enc.Encode(array)
	Throw(err, "")
}

func main() {
	err := os.MkdirAll("output", 0777)
	Throw(err, "")
	mapData := transData()
	outputMap(mapData)
	outputArray(mapData)
	log.Println("done")
}
