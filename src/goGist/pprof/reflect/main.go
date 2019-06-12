package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/pprof"
)

type People struct {
	Age int
}

func Foo(k interface{}) {
	_ = reflect.ValueOf(k)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}

		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 50000000; i++ {
		var k People
		Foo(k)
	}
}
