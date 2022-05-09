package main

import (
	"election-service/configs"
	"election-service/internal/core/models"
	"election-service/internal/core/services/candidatesrv"
	"election-service/internal/core/services/electionsrv"
	"election-service/internal/core/services/votersrv"
	"election-service/internal/handlers/candidatehdl"
	"election-service/internal/handlers/electionhdl"
	"election-service/internal/handlers/voterhdl"
	"election-service/internal/repositories/candidaterepo"
	"election-service/internal/repositories/candidatevoterepo"
	"election-service/internal/repositories/electionrepo"
	"election-service/internal/repositories/voterrepo"
	"election-service/internal/routes"
	"election-service/pkg"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	load()
	pkg.Init()

	db := pkg.InitDB()
	rest := pkg.InitRest()
	db.AutoMigrate(models.CandidateVote{}, models.Candidate{}, models.Election{}, models.Voter{})

	electionRepository := electionrepo.NewSql(db)
	candidateRepository := candidaterepo.NewSql(db)
	candidateVoteRepository := candidatevoterepo.NewSql(db)
	voterRepository := voterrepo.NewSql(db)

	electionService := electionsrv.New(electionRepository, candidateRepository, candidateVoteRepository)
	candidateService := candidatesrv.New(candidateRepository, candidateVoteRepository)
	voterService := votersrv.New(voterRepository, electionRepository, candidateVoteRepository)

	electionHandler := electionhdl.NewRest(electionService)
	candidateHandler := candidatehdl.NewRest(candidateService)
	voterHandler := voterhdl.NewRest(voterService)

	routes.ElectionEndpoints(rest, electionHandler)
	routes.CandidateEndpoints(rest, candidateHandler)
	routes.VoteEndpoints(rest, voterHandler)

	pkg.Started()
	rest.Listen(fmt.Sprintf(":%v", configs.Server.Port))
}

func load() {
	rand.Seed(time.Now().UnixNano())
	pkg.InitLog()
	configs.Load("config")
	pkg.Info("OO", "Configuration loaded")
	pkg.Info("__", "APP=%v", configs.App)
	pkg.Info("__", "DATABASE=%v", configs.Database)
	pkg.Info("__", "SERVER=%v", configs.Server)
}
