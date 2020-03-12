package main

import (
	"configmap-test/config"
	"github.com/sirupsen/logrus"
)

func main()  {
	logrus.Error(config.Conf.Kafka)
	logrus.Error(config.Conf.FactoryList)
}
