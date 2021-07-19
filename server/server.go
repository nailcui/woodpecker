package server

import (
	"sync"
	"woodpecker/single"
)

func Start(config *single.Config) {
	engine := single.NewSingleEngine(config)
	engine.Start()
	group := sync.WaitGroup{}
	group.Add(1)
	group.Wait()
}
