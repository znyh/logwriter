package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/logwriter/internal/model"
	"github.com/znyh/middle-end/library/pkg/kafka"
	logserver "github.com/znyh/proto/logserver"
)

func NewKafkaConsumer() (consumer *kafka.Consumer, cf func(), err error) {
	cfg := &kafka.Config{}
	if err = paladin.Watch("kafka.txt", cfg); err != nil {
		return nil, nil, err
	}
	log.Info("[NewKafka.Consumer] addrs = %v", cfg.Addr)
	consumer = &kafka.Consumer{}
	*consumer, err = kafka.NewConsumer(cfg.Addr, model.APPID)
	if err != nil {
		log.Error("new kafka consumer fail, msg: %v", err)
		return
	}
	cf = func() { _ = consumer.Close() }
	return
}

func (d *dao) SubKafka() (err error) {

	err = d.consumer.Consume(map[string]kafka.Handler{
		logserver.LogServerTopic: kafka.Handler{
			Run: func(msg []byte, args ...interface{}) {
				log.Info("=====>data:%+v", string(msg))
			},
			Args: nil,
		},
	})

	return
}
