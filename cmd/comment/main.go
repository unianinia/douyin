package main

import (
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/dal"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/constants"
	"douyin/pkg/mw"
	tracer "douyin/pkg/trace"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"time"
)

const serviceAddr = "127.0.0.1:9097"

func main() {
	dal.Init()
	rpc.InitRPC()
	tracer.InitJaeger(serviceAddr)

	r, err := etcd.NewEtcdRegistry([]string{
		constants.EtcdAddress,
	})

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)

	svr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), // server name
		server.WithMiddleware(mw.CommonMiddleware),                                                        // middleware
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithRegistry(r),                                             // registry
		server.WithReadWriteTimeout(10*time.Minute),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
