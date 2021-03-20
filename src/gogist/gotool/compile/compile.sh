#!/usr/bin/env bash
# reference: https://colobu.com/2018/12/29/get-assembly-output-for-go-programs/

# method-1
go tool compile -N -l -S main.go

# method-2
# go build -gcflags -S main.go

# method-3
# go tool compile -N -l once.go && go tool objdump once.o