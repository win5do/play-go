package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"playGo/src/gogist/rpc/buitinRpc/common"
)

type Arith int

func (t *Arith) Multiply(args *common.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *common.Args, quo *common.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	if err := rpc.Register(arith); err != nil {
		log.Fatal("register error:", err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	if err := http.Serve(l, nil); err != nil {
		log.Fatal("serve error:", err)
	}
}
