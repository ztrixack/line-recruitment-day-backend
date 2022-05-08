package main

import (
	"election-service/configs"
	"election-service/internal/core/models"
	"election-service/pkg"
	"math/rand"
	"time"
)

func main() {
	load()
	forever := make(chan bool)
	pkg.Init()

	db := pkg.InitDB()
	db.AutoMigrate(models.CandidateVote{}, models.Candidate{}, models.Election{}, models.Voter{})

	_ = db

	pkg.Started()
	<-forever
}

func load() {
	rand.Seed(time.Now().UnixNano())
	pkg.InitLog()
	configs.Load("config")
	pkg.Info("OO", "Configuration loaded")
	pkg.Info("__", "APP=%v", configs.App)
	pkg.Info("__", "DATABASE=%v", configs.Database)
}
