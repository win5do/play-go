package main

import (
	"flag"
	"log"
	"os"
)

var (
	Version   string
	Branch    string
	Commit    string
	BuildTime string
	lowercase string // 小写也可以
)

func main() {
	versionFlag := flag.Bool("version", false, "print the version")
	flag.Parse()

	if *versionFlag {
		log.Printf("Version: %s\n", Version)
		log.Printf("Branch: %s\n", Branch)
		log.Printf("Commit: %s\n", Commit)
		log.Printf("BuildTime: %s\n", BuildTime)
		log.Printf("lowercase: %s\n", lowercase)
		os.Exit(0)
	}

	log.Println("run main.go")
}
