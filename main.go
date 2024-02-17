package main

import (
	"fmt"
	"log"
	"offerBook/config"
	"offerBook/dao/redis"
	"offerBook/initialize"
	"offerBook/router"
)

func main() {
	err := initialize.Init()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err = redis.DB.Close(); err != nil {
			log.Println(err)
		}
	}()

	r := router.Init()
	addr := fmt.Sprintf("%s:%d", config.Cfg.Server.Addr, config.Cfg.Server.Port)
	log.Panic(r.Run(addr))
}
