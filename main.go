package main

import (
	"configmap-test/config"
	"github.com/sirupsen/logrus"
	"time"
)

func main()  {
	logrus.Info(config.GetSpecific().Kafka)
	logrus.Info(config.GetSpecific().FactoryList)
	t1 := time.NewTimer(time.Hour * 1)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Minute * 1)
			logrus.Info(config.GetSpecific().FactoryList)
		}
	}
}
