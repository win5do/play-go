package main

import (
	"fmt"
	"sync"
	"time"
)

/*
pool创建的时候是不能指定大小的，所有sync.Pool的缓存对象数量是没有限制的（只受限于内存），因此使用sync.pool是没办法做到控制缓存对象数量的个数的。

源码src/pkg/sync/pool.go在init的时候注册了一个poolCleanup函数，它会清除所有的pool里面的所有缓存的对象，该函数注册进去之后会在每次gc之前都会调用，因此sync.Pool缓存的期限只是两次gc之间这段时间。正因gc的时候会清掉缓存对象，也不用担心pool会无限增大的问题。

sync.Pool的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少gc的负担。
*/

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return b
	},
}

func main() {
	a := time.Now().Unix()
	// 不使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	// 使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().([]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}
