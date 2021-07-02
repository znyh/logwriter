package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/logwriter/internal/di"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("logwriter start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("logwriter exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
