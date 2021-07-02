package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// Ping ping the resource.
func (s *Service) Ping2(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}
