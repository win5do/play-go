#!/usr/bin/env bash
go build main.go

./main -cpuprofile=main.prof

go tool pprof -http=":8081" main main.prof