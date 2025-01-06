/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2025-01-06 15:57:36
# File Name: server_mgr_test.go
# Description:
####################################################################### */

package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ant-libs-go/config"
	"github.com/ant-libs-go/config/options"
	"github.com/ant-libs-go/config/parser"
	"github.com/ant-libs-go/rpcx/pb"
)

type Args struct {
	A int
	B int
}

func (t *Args) GetHeader() *pb.Header {
	return nil
}

type Reply struct {
	C int
}

func (t *Reply) GetHeader() *pb.Header {
	return nil
}

type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Println("---")
	reply.C = args.A * args.B
	return nil
}

func TestAA(m *testing.T) {
	config.NewConfig(parser.NewTomlParser(),
		options.WithCheckInterval(10),
		options.WithCfgSource("test.toml"),
		options.WithOnErrorFn(func(e error) { fmt.Println(e) }))

	err := StartServer("demo_serverzk", new(Arith))
	fmt.Println(err)
	// fmt.Printf("App rpcx listening on %s\n", RpcxServer("task_server").Address().String())

	time.Sleep(time.Hour)
}

func TestBB(m *testing.T) {
	config.NewConfig(parser.NewTomlParser(),
		options.WithCheckInterval(10),
		options.WithCfgSource("test.toml"),
		options.WithOnErrorFn(func(e error) { fmt.Println(e) }))

	err := StartServer("demo_servernc", new(Arith))
	fmt.Println(err)
	fmt.Printf("App rpcx listening on %s\n", RpcxServer("task_server").Address().String())

	time.Sleep(time.Hour)
}
