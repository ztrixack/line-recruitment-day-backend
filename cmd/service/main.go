package main

import (
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
}
