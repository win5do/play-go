package main

import (
	"io/ioutil"
	"regexp"
	"testing"
	"archive/tar"
	"os"
)

func TestHttp(t *testing.T) {
	bodyByte := tryHttp(3)
	if len(bodyByte) == 0 {
		t.Error("get失败")
	}
}

func TestMatch(t *testing.T) {
	htmlTar, err := os.Open("address_test.html.tar.gz")
	tr := tar.NewReader(htmlTar)
	tr.Read(b)
	if err != nil {
		t.Fatal("打开测试html失败")
	}
	body := string(htmlByte)
	reg := regexp.MustCompile(`<p class="MsoNormal">.*?<span lang="EN-US">(\d+)<span>.+?<span style="font-family: 宋体">[\p{Zs}\s]*(\p{Han}+)</span>.*?</p>`)
	matched := reg.FindAllStringSubmatch(body, -1)
	if len(matched) != 3508 {
		t.Error("匹配不全")
	}
}

func mockMapData(n int) map[string]*Address {
	mapData := map[string]*Address{}

	if n <= 0 {
		return mapData
	}

	item := Address{
		Value:    "100000",
		Label:    "北京市",
		Children: map[string]*Address{},
	}

	mapData[string(0)] = &item
	item_1 := item
	item_1.Children = mockMapData(n - 2)
	mapData[string(1)] = &item_1
	item_2 := item
	item_2.Children = mockMapData(n - 1)
	mapData[string(2)] = &item_2
	return mapData
}

func TestMapToArray(t *testing.T) {
	mapData := mockMapData(3)
	array := mapToArray(mapData)
	if (array[2].Children[2].Children[2].Value != "100000") {
		t.Error("结构有误")
	}
}
