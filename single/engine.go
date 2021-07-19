package single

import (
	"woodpecker/checker"
	"woodpecker/core"
	"woodpecker/logger"
	"woodpecker/notifier"
)

type Engine struct {
	Config          *Config
	ResourceFileMap map[string]*core.ResourceFile
	CheckerMap      map[string]checker.Checker
	NotifierMap     map[string]notifier.Notifier
}

func NewSingleEngine(config *Config) core.Engine {
	return &Engine{
		Config:          config,
		ResourceFileMap: map[string]*core.ResourceFile{},
		CheckerMap:      map[string]checker.Checker{},
		NotifierMap:     map[string]notifier.Notifier{},
	}
}

func (e *Engine) Start() {
	err := LoadAllResourceFile(e)
	if err != nil {
		logger.Error("load resource file error")
		panic(err)
	}
	err = LoadAllResource(e)
	if err != nil {
		logger.Error("load resource error")
		panic(err)
	}
	logger.Info("load all resource success")
	err = EnableAllChecker(e)
}
