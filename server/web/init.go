package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"phanes/config"
	"phanes/server/web/middleware"
	"phanes/server/web/v1"
	"phanes/utils"
	"time"
)

var (
	webName           string
	webAddr           string
	defaultListenAddr = ":7771"
	srv               server.Server
)

func Init() micro.Option {
	webName = config.Conf.Name + "-http"

	if config.Conf.Http.HttpListen != "" {
		defaultListenAddr = config.Conf.Http.HttpListen
	}

	srv = http.NewServer(
		server.Name(webName),
		server.Version(config.Conf.Version),
		server.Address(defaultListenAddr),
		server.RegisterTTL(time.Second*30),
		server.RegisterInterval(time.Second*15),
		server.Registry(etcd.NewRegistry(registry.Addrs(config.EtcdAddr))),
	)

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// register routers
	v1Group := router.Group("v1", middleware.OtelMiddleware())
	v1.Init(v1Group)

	utils.Throw(srv.Handle(srv.NewHandler(router)))
	utils.Throw(srv.Start())
	if config.Conf.Traefik.Enabled {
		utils.Throw(config.Register(webName, srv))
	}

	return micro.Server(srv)
}
