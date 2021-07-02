package http

import (
	"net/http"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	pb "github.com/znyh/proto/logwriter"
)

var svc pb.LogwriterServer

// New new a bm server.
func New(s pb.LogwriterServer) (engine *bm.Engine, err error) {
	var cfg struct {
		Server *bm.ServerConfig
		ct     paladin.TOML
	}
	if err = paladin.Get("http.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(cfg.Server)
	pb.RegisterLogwriterBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/logwriter")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	c.JSON("Golang 大法好 !!!", nil)
}
