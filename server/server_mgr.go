/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-10-30 22:01:02
# File Name: server_mgr.go
# Description:
####################################################################### */

package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ant-libs-go/config"
	"github.com/ant-libs-go/safe_stop"
	rpcx_server "github.com/smallnest/rpcx/server"
	"google.golang.org/grpc"
)

var (
	once    sync.Once
	lock    sync.RWMutex
	servers map[string]*srv
)

func init() {
	servers = map[string]*srv{}
}

type rpcxConfig struct {
	Rpcx *struct {
		Cfgs map[string]*Cfg `toml:"server"`
	} `toml:"rpcx"`
}

type Cfg struct {
	// dial
	DialAddr         string        `toml:"addr"`
	DialReadTimeout  time.Duration `toml:"read_timeout"`
	DialWriteTimeout time.Duration `toml:"write_timeout"`

	EnableGrpc bool `toml:"enable_grpc"`

	// register
	Register string `toml:"register"`

	// zookeeper register
	RegisterZkServers        []string      `toml:"register_zk_servers"`
	RegisterZkBasePath       string        `toml:"register_zk_basepath"`
	RegisterZkUpdateInterval time.Duration `toml:"register_zk_update_interval"`

	// nacos register
	RegisterNcServers     []string `toml:"register_nc_servers"`
	RegisterNcNamespaceId string   `toml:"register_nc_namespace_id"`
	RegisterNcCacheDir    string   `toml:"register_nc_cache_dir"`
	RegisterNcLogDir      string   `toml:"register_nc_log_dir"`
	RegisterNcLogLevel    string   `toml:"register_nc_log_level"`
	RegisterNcAccessKey   string   `toml:"register_nc_access_key"`
	RegisterNcSecretKey   string   `toml:"register_nc_secret_key"`
}

func RpcxServer(name string) (r *rpcx_server.Server) {
	return Server(name).s
}

func StartServerAndGrpc(name string, rcvr interface{}, grcvr func(*grpc.Server)) (err error) {
	safe_stop.Lock(1)
	var srv *srv
	if srv, err = SafeServer(name); err == nil {
		err = srv.StartAndGrpc(name, rcvr, grcvr)
	}
	return
}

func StartServer(name string, rcvr interface{}) (err error) {
	safe_stop.Lock(1)
	var srv *srv
	if srv, err = SafeServer(name); err == nil {
		err = srv.Start(name, rcvr)
	}
	return
}

func StopServer(name string) (err error) {
	defer safe_stop.Unlock()
	var srv *srv
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if srv, err = SafeServer(name); err == nil {
		err = srv.Shutdown(ctx)
	}
	return
}

func Server(name string) (r *srv) {
	var err error
	if r, err = getServer(name); err != nil {
		panic(err)
	}
	return
}

func SafeServer(name string) (r *srv, err error) {
	return getServer(name)
}

func getServer(name string) (r *srv, err error) {
	lock.RLock()
	r = servers[name]
	lock.RUnlock()
	if r == nil {
		r, err = addServer(name)
	}
	return
}

func addServer(name string) (r *srv, err error) {
	var cfg *Cfg
	if cfg, err = loadCfg(name); err != nil {
		return
	}
	if r, err = NewRpcxServer(cfg); err != nil {
		return
	}

	lock.Lock()
	servers[name] = r
	lock.Unlock()
	return
}

func loadCfg(name string) (r *Cfg, err error) {
	var cfgs map[string]*Cfg
	if cfgs, err = loadCfgs(); err != nil {
		return
	}
	if r = cfgs[name]; r == nil {
		err = fmt.Errorf("rpcx#%s not configed", name)
		return
	}
	return
}

func loadCfgs() (r map[string]*Cfg, err error) {
	r = map[string]*Cfg{}

	cfg := config.Get(&rpcxConfig{}).(*rpcxConfig)
	if err == nil && (cfg.Rpcx == nil || cfg.Rpcx.Cfgs == nil || len(cfg.Rpcx.Cfgs) == 0) {
		err = fmt.Errorf("not configed")
	}
	if err != nil {
		err = fmt.Errorf("rpcx load cfgs error, %s", err)
		return
	}
	r = cfg.Rpcx.Cfgs
	return
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
