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
	"net"
	"os"
	"strings"
	"time"

	"github.com/ant-libs-go/rpcx"
	"github.com/ant-libs-go/rpcx/pb"
	"github.com/ant-libs-go/util"
	"github.com/ant-libs-go/util/logs"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/rcrowley/go-metrics"
	nc_plugin "github.com/rpcxio/rpcx-nacos/serverplugin"
	zk_plugin "github.com/rpcxio/rpcx-zookeeper/serverplugin"
	"github.com/rpcxio/rpcxplus/grpcx"
	uuid "github.com/satori/go.uuid"
	rpcx_server "github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"google.golang.org/grpc"
)

var pwd, _ = os.Getwd()

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
	resp.Header = rpcx.BuildRpcxHeader(req.Header.TraceId, "")
	log.Infof("Run success: %+v", resp)
	return
}

type srv struct {
	s        *rpcx_server.Server
	gs       *grpcx.GrpcServerPlugin
	register rpcx_server.RegisterPlugin
	isServe  bool
	cfg      *Cfg
	rcvrs    map[string]interface{}
}

func (this *srv) StartAndGrpc(name string, rcvr interface{}, grcvr func(*grpc.Server)) (err error) {
	if this.gs != nil {
		this.gs.RegisterService(grcvr)
	}
	return this.Start(name, rcvr)
}

func (this *srv) Start(name string, rcvr interface{}) (err error) {
	this.rcvrs[name] = rcvr
	for _, rcvr := range this.rcvrs {
		err = this.s.RegisterName(name, rcvr, "")
		if err != nil {
			return
		}
	}
	if this.isServe == true {
		return
	}
	this.isServe = true

	if this.gs != nil {
		go func() { err = this.gs.Start() }()
	}
	time.Sleep(time.Second)
	if err == nil {
		go func() {
			if err = this.s.Serve("tcp", this.cfg.DialAddr); err == rpcx_server.ErrServerClosed {
				err = nil
			}
		}()
	}
	time.Sleep(time.Second)
	return
}

func (this *srv) Shutdown(ctx context.Context) (err error) {
	err = this.s.Shutdown(ctx)
	if err == nil && this.gs != nil {
		err = this.gs.Close()
	}
	if err == nil {
		switch v := this.register.(type) {
		case *zk_plugin.ZooKeeperRegisterPlugin:
			err = v.Stop()
		case *nc_plugin.NacosRegisterPlugin:
			err = v.Stop()
		}
	}
	return
}

func NewRpcxServer(cfg *Cfg) (r *srv, err error) {
	r = &srv{
		s:     rpcx_server.NewServer(buildServerOptions(cfg)...),
		cfg:   cfg,
		rcvrs: map[string]interface{}{},
	}

	if cfg.EnableGrpc == true {
		r.gs = grpcx.NewGrpcServerPlugin()
		r.s.Plugins.Add(r.gs)
	}

	if len(cfg.Register) > 0 {
		var lisIp, lisPort string
		if lisIp, err = util.GetLocalIP(); err != nil || len(lisIp) == 0 {
			err = fmt.Errorf("ip parse fail, %s", err)
		}
		if err == nil {
			if t := strings.Split(cfg.DialAddr, ":"); len(t) != 2 || len(t[1]) == 0 {
				err = fmt.Errorf("port parse fail")
			} else {
				lisPort = t[1]
			}
		}
		if err != nil {
			return
		}

		switch cfg.Register {
		case "zookeeper_register":
			updateInterval := time.Minute
			if cfg.RegisterZkUpdateInterval > 0 {
				updateInterval = cfg.RegisterZkUpdateInterval * time.Millisecond
			}
			r.register = &zk_plugin.ZooKeeperRegisterPlugin{
				ServiceAddress:   fmt.Sprintf("tcp@%s:%s", lisIp, lisPort),
				ZooKeeperServers: cfg.RegisterZkServers,
				BasePath:         cfg.RegisterZkBasePath, // todo
				Metrics:          metrics.NewRegistry(),
				UpdateInterval:   updateInterval,
			}
			err = r.register.(*zk_plugin.ZooKeeperRegisterPlugin).Start()
		case "nacos_register":
			var host, port string
			sc := []constant.ServerConfig{}
			for _, addr := range cfg.RegisterNcServers {
				if host, port, err = net.SplitHostPort(addr); err != nil {
					err = fmt.Errorf("server addr parse fail, %s", err)
					return
				}
				sc = append(sc, *constant.NewServerConfig(host, uint64(util.StrToInt64(port, 80))))
			}
			cc := constant.NewClientConfig(
				constant.WithTimeoutMs(5000),
				constant.WithNamespaceId(cfg.RegisterNcNamespaceId),
				constant.WithNotLoadCacheAtStart(true),
				constant.WithCacheDir(util.AbsPath(cfg.RegisterNcCacheDir, pwd)),
				constant.WithLogDir(util.AbsPath(cfg.RegisterNcLogDir, pwd)),
				constant.WithLogLevel(cfg.RegisterNcLogLevel),
				constant.WithAccessKey(cfg.RegisterNcAccessKey),
				constant.WithSecretKey(cfg.RegisterNcSecretKey),
			)
			r.register = &nc_plugin.NacosRegisterPlugin{
				ServiceAddress: fmt.Sprintf("tcp@%s:%s", lisIp, lisPort),
				ClientConfig:   *cc,
				ServerConfig:   sc,
				// Cluster: "",     // cfg.RegisterNcBasePath
				// Group: "aaaa",
			}
			err = r.register.(*nc_plugin.NacosRegisterPlugin).Start()
		}

		if r.register != nil {
			r.s.Plugins.Add(r.register)
			r.s.Plugins.Add(&serverplugin.MetricsPlugin{Registry: metrics.DefaultRegistry, Prefix: "rpcx."})
		}
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
