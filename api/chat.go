package main

import (
	"flag"
	"fmt"
	"log"
	"zero-chat/api/internal/common/imserver"
	"zero-chat/api/internal/config"
	"zero-chat/api/internal/handler"
	"zero-chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//ct, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	//defer cancel()
	//go tmp(ct, ctx)

	imServer, err := imserver.NewImServer(ctx.Redis, ctx.KafkaReader)
	log.Printf("imServer:s%", imServer)
	if err != nil {
		log.Fatal(err)
	}
	go imServer.SubscribeTwo()
	//go imServer.Subscribe()
	go imServer.Run()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
