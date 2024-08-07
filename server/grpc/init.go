package grpc

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"phanes/config"
	"phanes/server/grpc/middleware"
	"time"
)

func Init() micro.Option {

	opts := []server.Option{
		server.Name(config.Conf.Name + "-grpc"),
		server.Version(config.Conf.Version),
		server.RegisterTTL(time.Second * 30),
		server.RegisterInterval(time.Second * 15),
		server.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddr))),
		server.WrapHandler(middleware.ServerTraceWrapper()),
		server.WrapHandler(middleware.Log()),
	}

	if config.Conf.Grpc.GrpcListen != "" {
		opts = append(opts, server.Address(config.Conf.Grpc.GrpcListen))
	}

	if config.Conf.Grpc.DiscoveryListen != "" {
		opts = append(opts, server.Advertise(config.Conf.Grpc.DiscoveryListen))
	}

	srv := grpc.NewServer(opts...)
	// register grpc services
	// example: utils.Throw(micro.RegisterHandler(srv, new(App)))

	// ⚠️Waring!!!: Your service struct Name Must seem to the .proto file service Name
	// utils.Throw(micro.RegisterHandler(srv, new(v1.User)))

	return micro.Server(srv)
}
