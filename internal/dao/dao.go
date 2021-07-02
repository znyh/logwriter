package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/google/wire"
	"github.com/znyh/library/pkg/kafka"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewKafkaConsumer)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	SubKafka() (err error)
}

// dao dao.
type dao struct {
	db       *sql.DB
	redis    *redis.Redis
	consumer *kafka.Consumer
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB, c *kafka.Consumer) (d Dao, cf func(), err error) {
	d = &dao{
		db:       db,
		redis:    r,
		consumer: c,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	_ = d.redis.Close()
	_ = d.db.Close()
	_ = d.consumer.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
