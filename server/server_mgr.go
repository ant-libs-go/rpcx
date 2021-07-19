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

	// register, only support zookeeper
	RegisterServers        []string      `toml:"register_servers"`
	RegisterBasePath       string        `toml:"register_base_path"`
	RegisterUpdateInterval time.Duration `toml:"register_update_interval"`
}

func StartDefaultServer(rcvr interface{}) (err error) {
	return StartServer("default", rcvr)
}

func StopDefaultServer() (err error) {
	return StopServer("default")
}

func DefaultServer() (r *rpcx_server.Server) {
	return Server("default").s
}

func StartServer(name string, rcvr interface{}) (err error) {
	safe_stop.Lock(1)
	var srv *srv
	if srv, err = SafeServer(name); err == nil {
		err = srv.Start("Server", rcvr)
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

	cfg := &rpcxConfig{}
	once.Do(func() {
		_, err = config.Load(cfg)
	})

	cfg = config.Get(cfg).(*rpcxConfig)
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
