/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2025-01-06 18:23:33
# File Name: client_mgr_test.go
# Description:
####################################################################### */

package client

import (
	"context"
	"fmt"
	"testing"

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
	reply.C = args.A * args.B
	return nil
}

func TestMain(m *testing.M) {
	config.NewConfig(parser.NewTomlParser(),
		options.WithCheckInterval(10),
		options.WithCfgSource("test.toml"),
		options.WithOnErrorFn(func(e error) { fmt.Println(e) }))

	req := &Args{A: 10, B: 20}
	resp := &Reply{}

	err := Call(context.TODO(), "demo_serverzk", "Mul", req, resp)
	fmt.Println(err)
	fmt.Println(resp)
}
