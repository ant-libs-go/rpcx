/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-08-12 14:50:38
# File Name: client.go
# Description:
####################################################################### */

package client

import (
	"time"

	"github.com/ant-libs-go/util"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

func NewRpcxClientPool(cfg *Cfg) (r *client.XClientPool, err error) {
	maxActive := util.If(cfg.PoolMaxActive > 0, cfg.PoolMaxActive, 10).(int)
	util.IfDo(len(cfg.DialServerName) == 0, func() { cfg.DialServerName = "Server" })
	discovery, err := buildDialDiscovery(cfg)
	if err != nil {
		return
	}
	r = client.NewXClientPool(maxActive, cfg.DialServerName,
		buildDialFailover(cfg), buildDialSelector(cfg), discovery, buildDialOption(cfg))
	return
}

// 暂只对MultipleServersDiscovery进行支持
func buildDialDiscovery(cfg *Cfg) (r client.ServiceDiscovery, err error) {
	switch cfg.DialDiscovery {
	case "multiple_servers_discovery":
		kvs := []*client.KVPair{}
		for _, addr := range cfg.DialAddrs {
			kvs = append(kvs, &client.KVPair{Key: addr})
		}
		r, err = client.NewMultipleServersDiscovery(kvs)
	case "zookeeper_discovery":
		r, err = client.NewZookeeperDiscovery(cfg.DiscoveryBasePath, cfg.DialServerName, cfg.DialAddrs, nil)
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
