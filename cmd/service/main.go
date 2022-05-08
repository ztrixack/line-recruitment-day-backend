package main

import (
	"election-service/configs"
	"election-service/pkg"
	"math/rand"
	"time"
)

func main() {
	load()
	forever := make(chan bool)
	<-forever
}

func load() {
	rand.Seed(time.Now().UnixNano())
	pkg.Init()
	pkg.InitLog()
	configs.Load("config")
	pkg.Info("OO", "Configuration loaded")
	pkg.Info("__", "APP=%v", configs.App)
}
