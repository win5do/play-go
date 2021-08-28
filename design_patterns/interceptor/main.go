package main

import (
	"fmt"

	"playGo/design_patterns/interceptor/account"
	"playGo/design_patterns/interceptor/proxy"
)

func main() {
	id := "100111"
	a := account.New(id, "ZhangSan", 100)
	// 拦截器
	a.Query(id)
	a.Update(id, 500)

	// 装饰器
	decoratorQuery(a)(id)
}

func init() {
	account.New = func(id, name string, value int) account.Account {
		a := &account.AccountImpl{id, name, value}
		p := &proxy.Proxy{a}
		return p
	}
}

func decoratorQuery(a account.Account) func(id string) int {
	fn := func(id string) int {
		fmt.Println("before")
		r := a.Query(id)
		fmt.Println("before")
		return r
	}
	return fn
}
