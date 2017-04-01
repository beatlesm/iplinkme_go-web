package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}
type Math int

func (m *Math) Multipy(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Quotient struct {
	Quo, Rem int
}

func (m *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	math := new(Math)
	rpc.Register(math)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn error", err)
		}
		jsonrpc.ServeConn(conn)
		//rpc.ServeConn(conn)
	}
}

//
