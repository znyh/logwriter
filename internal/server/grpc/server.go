package grpc

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	pb "github.com/znyh/proto/logwriter"
)

// New new a grpc server.
func New(svc pb.LogwriterServer) (ws *warden.Server, err error) {
	var cfg struct {
		Server *warden.ServerConfig
	}
	if err = paladin.Get("grpc.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(cfg.Server)
	pb.RegisterLogwriterServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
