/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-10-30 22:22:32
# File Name: server.go
# Description:
####################################################################### */

package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ant-libs-go/rpcx"
	"github.com/ant-libs-go/rpcx/pb"
	"github.com/ant-libs-go/util"
	"github.com/ant-libs-go/util/logs"
	"github.com/rcrowley/go-metrics"
	zookeeper_plugin "github.com/rpcxio/rpcx-zookeeper/serverplugin"
	uuid "github.com/satori/go.uuid"
	rpcx_server "github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type ServiceImpl struct{}

func (this *ServiceImpl) Before(header *pb.Header) (log *logs.SessLog) {
	if len(header.TraceId) == 0 {
		header.TraceId = uuid.NewV4().String()
	}
	log = logs.New(header.TraceId)
	return
}

func (this *ServiceImpl) Ping(ctx context.Context, req *pb.Ping_Req, resp *pb.Ping_Resp) (err error) {
	log := this.Before(req.Header)
	log.Infof("Request type: Ping, req: %+v", req)
	resp.Header = rpcx.BuildRpcxHeader(req.Header.TraceId)
	log.Infof("Run success: %+v", resp)
	return
}

type srv struct {
	s        *rpcx_server.Server
	register *zookeeper_plugin.ZooKeeperRegisterPlugin
	isServe  bool
	cfg      *Cfg
	rcvrs    map[string]interface{}
}

func (this *srv) Start(name string, rcvr interface{}) (err error) {
	this.rcvrs[name] = rcvr
	for _, rcvr := range this.rcvrs {
		err = this.s.RegisterName(name, rcvr, "")
		if err != nil {
			return
		}
	}

	go func() {
		if this.isServe == true {
			return
		}
		this.isServe = true

		if err = this.s.Serve("tcp", this.cfg.DialAddr); err == rpcx_server.ErrServerClosed {
			err = nil
		}
		if err == nil && this.register != nil {
			err = this.register.Start()
		}
	}()
	time.Sleep(time.Second)
	return
}

func (this *srv) Shutdown(ctx context.Context) (err error) {
	err = this.s.Shutdown(ctx)
	if err == nil && this.register != nil {
		err = this.register.Stop()
	}
	return
}

func NewRpcxServer(cfg *Cfg) (r *srv, err error) {
	r = &srv{
		s:     rpcx_server.NewServer(buildServerOptions(cfg)...),
		cfg:   cfg,
		rcvrs: map[string]interface{}{}}

	if len(cfg.RegisterServers) > 0 {
		var ip, port string
		if ip, err = util.GetLocalIP(); err != nil || len(ip) == 0 {
			err = fmt.Errorf("ip parse fail, %s", err)
		}
		if err == nil {
			if t := strings.Split(cfg.DialAddr, ":"); len(t) != 2 || len(t[1]) == 0 {
				err = fmt.Errorf("port parse fail")
			} else {
				port = t[1]
			}
		}
		if err != nil {
			return
		}

		r.register = &zookeeper_plugin.ZooKeeperRegisterPlugin{
			ServiceAddress:   fmt.Sprintf("tcp@%s:%s", ip, port),
			ZooKeeperServers: cfg.RegisterServers,
			BasePath:         cfg.RegisterBasePath,
			Metrics:          metrics.NewRegistry(),
			UpdateInterval:   time.Minute}
		if cfg.RegisterUpdateInterval > 0 {
			r.register.UpdateInterval = cfg.RegisterUpdateInterval * time.Millisecond
		}

		r.s.Plugins.Add(r.register)
		r.s.Plugins.Add(&serverplugin.MetricsPlugin{Registry: metrics.DefaultRegistry, Prefix: "rpcx."})
	}
	return
}

func buildServerOptions(cfg *Cfg) (r []rpcx_server.OptionFn) {
	if cfg.DialReadTimeout > 0 {
		r = append(r, rpcx_server.WithReadTimeout(cfg.DialReadTimeout*time.Millisecond))
	}
	if cfg.DialWriteTimeout > 0 {
		r = append(r, rpcx_server.WithWriteTimeout(cfg.DialWriteTimeout*time.Millisecond))
	}
	return
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
