/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-10-30 22:24:56
# File Name: rpcx.go
# Description:
####################################################################### */

package rpcx

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ant-libs-go/rpcx/pb"
	uuid "github.com/satori/go.uuid"
)

var (
	host       string
	servername string
)

func init() {
	host, _ = os.Hostname()
	servername = filepath.Base(os.Args[0])
}

func BuildRpcxHeader(traceId string) *pb.Header {
	if len(traceId) == 0 {
		traceId = uuid.NewV4().String()
	}
	return &pb.Header{
		Requester: servername + "#" + host,
		TraceId:   traceId,
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
		Metadata:  map[string]string{}}
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
