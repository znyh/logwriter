package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var cfg struct {
		Client *redis.Config
	}
	if err = paladin.Get("redis.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewRedis(cfg.Client)
	cf = func() { _ = r.Close() }
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func (d *dao) RedisSUBSCRIBE(ctx context.Context, topic string, cb func(data []byte)) (err error) {
	go func() {
		c, _ := redis.Dial("tcp", ":6379")
		//c := d.redis.Conn(context.Background())
		if c == nil {
			return
		}
		defer c.Close()
		psc := redis.PubSubConn{Conn: c}
		err = psc.Subscribe(topic)
		if err != nil {
			return
		}
		defer psc.Unsubscribe(topic)
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Info("%s: message: %s", v.Channel, v.Data)
				cb(v.Data)
			case redis.Subscription:
				log.Info("%s: %s, %d", v.Channel, v.Kind, v.Count)
			case error:
				log.Error("error:%+v", v)
				return
			}
		}
	}()

	return
}

func (d *dao) RedisPublish(ctx context.Context, topic, data string) (err error) {
	_, err = d.redis.Do(ctx, "PUBLISH", topic, data)
	return
}
