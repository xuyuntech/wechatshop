package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"errors"
	"github.com/xuyuntech/wechatshop/pkg/config"
)

func main() {
	configFilePath := os.Args[1]
	if len(configFilePath) == 0 {
		logrus.Fatal(errors.New("no config file provided"))
	}

	yConfig, err := config.LoadConfig(configFilePath)
	if err != nil {
		logrus.Fatal(err)
	}

	if yConfig.Server.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}