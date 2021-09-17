package single

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"woodpecker/logger"
)

func EnableAllChecker(engine *Engine) error {
	c := cron.New(cron.WithSeconds())
	c.Start()
	for name := range engine.CheckerMap {
		checker := engine.CheckerMap[name]
		if !checker.Enabled() {
			logger.Info("checker disabled", checker.GetName())
			continue
		}
		checker.Init()
		_, err := c.AddFunc(checker.GetCron(), func() {
			logger.Info(fmt.Sprintf("%s check start", checker.GetName()))
			if err := checker.Check(); err != nil {
				logger.Info(fmt.Sprintf("%s check unhealth error: %s\n", checker.GetName(), err.Error()))
				err = engine.NotifierMap[checker.GetNotifier()].Send(err.Error())
				if err != nil {
					logger.Info(fmt.Sprintf("%s check unhealth send msg to %s error: %s\n", checker.GetName(), checker.GetNotifier(), err.Error()))
				}
			}
		})
		if err != nil {
			logger.Info(fmt.Sprintf("add cron error: %e\n", err))
			return err
		}
	}
	return nil
}
