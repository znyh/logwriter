package service

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"github.com/znyh/logwriter/internal/dao"
	pb "github.com/znyh/proto/logwriter"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.LogwriterServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.txt", s.ac)

	//go func() {
	//	if err := s.dao.SubKafka(); err != nil {
	//		log.Error("SubKafka error:%+v", err)
	//	}
	//}()

	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
