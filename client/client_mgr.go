/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-08-12 15:25:26
# File Name: client_mgr.go
# Description:
####################################################################### */

package client

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ant-libs-go/config"
	"github.com/ant-libs-go/config/options"
	"github.com/ant-libs-go/rpcx/pb"
	"github.com/smallnest/rpcx/client"
	"google.golang.org/protobuf/proto"
)

var (
	once  sync.Once
	lock  sync.RWMutex
	pools map[string]*client.XClientPool
)

func init() {
	pools = map[string]*client.XClientPool{}
}

// type Message interface {
//	proto.Message
//	GetHeader() *pb.Header
//}

type rpcxConfig struct {
	Rpcx *struct {
		Cfgs map[string]*Cfg `toml:"client"`
	} `toml:"rpcx"`
}

type Cfg struct {
	// dial
	DialAddrs          []string      `toml:"addrs"`
	DialDiscovery      string        `toml:"discovery"`
	DialFailMode       string        `toml:"fail_mode"`
	DialSelectMode     string        `toml:"select_mode"`
	DialConnectTimeout time.Duration `toml:"dial_timeout"`

	// zookeeper register
	DiscoveryZkBasePath string `toml:"discovery_zk_basepath"`

	// nacos register
	RegisterNcNamespaceId string `toml:"register_nc_namespace_id"`
	RegisterNcCacheDir    string `toml:"register_nc_cache_dir"`
	RegisterNcLogDir      string `toml:"register_nc_log_dir"`
	RegisterNcLogLevel    string `toml:"register_nc_log_level"`
	RegisterNcAccessKey   string `toml:"register_nc_access_key"`
	RegisterNcSecretKey   string `toml:"register_nc_secret_key"`

	// pool
	PoolMaxActive int `toml:"pool_max_active"` // 最大活跃连接数
}

// 验证Rpcx实例的配置正确性与连通性。
// 参数names是实例的名称列表，如果为空则检测所有配置的实例
func Valid(names ...string) (err error) {
	if len(names) == 0 {
		var cfgs map[string]*Cfg
		if cfgs, err = loadCfgs(); err != nil {
			return
		}
		for k := range cfgs {
			names = append(names, k)
		}
	}
	for _, name := range names {
		if err == nil {
			// 所有Server都必须对Ping方法进行实现
			err = Call(context.Background(), name, "Ping", &pb.Ping_Req{}, &pb.Ping_Resp{})
		}
		if err != nil {
			err = fmt.Errorf("rpcx#%s is invalid, %s", name, err)
			return
		}
	}
	return
}

func Call(ctx context.Context, name string, method string, req proto.Message, resp proto.Message) (err error) {
	// if req.GetHeader() == nil {
	//	fl := reflect.ValueOf(req).Elem().FieldByName("Header")
	//	// if fl.IsValid() && fl.CanSet() && fl.Type().String() == "*common.Header" {
	//	if fl.IsValid() && fl.CanSet() {
	//		fl.Set(reflect.ValueOf(rpcx.BuildRpcxHeader("", "")))
	//	}
	//}
	var cli client.XClient
	cli, err = SafeClient(name)
	if err == nil {
		err = cli.Call(ctx, method, req, resp)
	}
	return
}

func Client(name string) (r client.XClient) {
	r = Pool(name).Get()
	return
}

func SafeClient(name string) (r client.XClient, err error) {
	var pool *client.XClientPool
	pool, err = SafePool(name)
	if err != nil {
		return
	}
	r = pool.Get()
	return
}

func Pool(name string) (r *client.XClientPool) {
	var err error
	if r, err = getPool(name); err != nil {
		panic(err)
	}
	return
}

func SafePool(name string) (r *client.XClientPool, err error) {
	return getPool(name)
}

func getPool(name string) (r *client.XClientPool, err error) {
	lock.RLock()
	r = pools[name]
	lock.RUnlock()
	if r == nil {
		r, err = addPool(name)
	}
	return
}

func addPool(name string) (r *client.XClientPool, err error) {
	var cfg *Cfg
	if cfg, err = loadCfg(name); err != nil {
		return
	}
	r, err = NewRpcxClientPool(name, cfg)

	lock.Lock()
	pools[name] = r
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

	once.Do(func() {
		config.Get(&rpcxConfig{}, options.WithOpOnChangeFn(func(cfg interface{}) {
			lock.Lock()
			defer lock.Unlock()
			pools = map[string]*client.XClientPool{}
		}))
	})

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
