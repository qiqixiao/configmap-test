package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
var(
	Conf 			*Config
)

type iKafaConsumer struct {
	Bootstrap	string		`yaml:"bootstrap"`
	Topics		[]string	`yaml:"topics"`
	GroupID		string		`yaml:"group_id"`
}

type iFactory struct {
	ID		uint	`yaml:"id"`
	Topics	string	`yaml:"topics"`
}

type SpecificConf struct{
	Kafka 		iKafaConsumer		`yaml:"kafka_consumer"`
	FactoryList	map[string]iFactory	`yaml:"factory_list"`
}

type Config struct{
	AppName	string	`yaml:"appname"`
	Port	string	`yaml:"port"`
	Mode	string	`yaml:"mode"`
}


func GetSpecific() (s *SpecificConf) {
	path := "conf-" + Conf.Mode + ".yml"
	specific, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("specific yamlFile.Get err %v ", err)
		return nil
	}
	err = yaml.Unmarshal(specific, &s)
	if err != nil  {
		logrus.Error("specific yamlFile parse error: ", err)
		return nil
	}
	return s
}


func init() {
	Conf = &Config{}
	configFile, err := ioutil.ReadFile("./config/conf.yml")
	if err != nil {
		logrus.Errorf("yamlFile.Get err %v ", err)
	}
	err = yaml.Unmarshal(configFile, Conf)
	if err != nil  {
		logrus.Error("yamlFile parse error: ", err)
	}
}