/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-08-12 14:50:38
# File Name: client.go
# Description:
####################################################################### */

package client

import (
	"net"
	"os"
	"time"

	"github.com/ant-libs-go/util"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	nc_client "github.com/rpcxio/rpcx-nacos/client"
	zk_client "github.com/rpcxio/rpcx-zookeeper/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var pwd, _ = os.Getwd()

func NewRpcxClientPool(name string, cfg *Cfg) (r *client.XClientPool, err error) {
	maxActive := util.If(cfg.PoolMaxActive > 0, cfg.PoolMaxActive, 10).(int)
	discovery, err := buildDialDiscovery(name, cfg)
	if err != nil {
		return
	}
	r = client.NewXClientPool(maxActive, name,
		buildDialFailover(cfg), buildDialSelector(cfg), discovery, buildDialOption(cfg))
	return
}

// 暂只对MultipleServersDiscovery进行支持
func buildDialDiscovery(name string, cfg *Cfg) (r client.ServiceDiscovery, err error) {
	switch cfg.DialDiscovery {
	case "multiple_servers_discovery":
		kvs := []*client.KVPair{}
		for _, addr := range cfg.DialAddrs {
			kvs = append(kvs, &client.KVPair{Key: addr})
		}
		r, err = client.NewMultipleServersDiscovery(kvs)
	case "zookeeper_discovery":
		r, err = zk_client.NewZookeeperDiscovery(cfg.DiscoveryZkBasePath, name, cfg.DialAddrs, nil)
	case "nacos_discovery":
		var host, port string
		sc := []constant.ServerConfig{}
		for _, addr := range cfg.DialAddrs {
			if host, port, err = net.SplitHostPort(addr); err != nil {
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
		r, err = nc_client.NewNacosDiscovery(name, "", "", *cc, sc)
	}
	return
}

func buildDialFailover(cfg *Cfg) (r client.FailMode) {
	r = client.Failover
	switch cfg.DialFailMode {
	case "fast":
		r = client.Failfast
	case "over":
		r = client.Failover
	case "try":
		r = client.Failtry
	case "backup":
		r = client.Failbackup
	}
	return
}

func buildDialSelector(cfg *Cfg) (r client.SelectMode) {
	r = client.RandomSelect
	switch cfg.DialSelectMode {
	case "random":
		r = client.RandomSelect
	case "round_robin":
		r = client.RoundRobin
	case "weighted_round_robin":
		r = client.WeightedRoundRobin
	case "weighted_icmp":
		r = client.WeightedICMP
	case "consistent_hash":
		r = client.ConsistentHash
	case "closest":
		r = client.Closest
	case "select_by_user":
		r = client.SelectByUser
	}
	return
}

func buildDialOption(cfg *Cfg) (r client.Option) {
	r = client.DefaultOption
	r.SerializeType = protocol.ProtoBuffer
	if cfg.DialConnectTimeout > 0 {
		r.ConnectTimeout = cfg.DialConnectTimeout * time.Millisecond
	} else {
		r.ConnectTimeout = 1000 * time.Millisecond
	}
	return
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
