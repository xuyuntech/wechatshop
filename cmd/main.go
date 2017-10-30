package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"errors"
	"github.com/xuyuntech/wechatshop/pkg/config"
	"fmt"
	"github.com/xuyuntech/wechatshop/pkg/models"
	"github.com/xuyuntech/wechatshop/pkg/manager"
	"github.com/xuyuntech/wechatshop/pkg/api"
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

	engine, err := models.NewEngine(models.DBOptions{
		User: yConfig.DB.User,
		Password: yConfig.DB.Password,
		Name: yConfig.DB.Name,
		Host: yConfig.DB.Host,
		Port: yConfig.DB.Port,
	}, models.Tables())
	if err != nil {
		logrus.Fatal(fmt.Errorf("new db engine err: %v", err))
	}

	controllerManager, err := manager.NewDefaultManager(yConfig, engine)
	if err != nil {
		logrus.Fatal(fmt.Errorf("new controller manager err: %v", err))
	}

	appApi, err := api.NewApi(&api.ApiConfig{
		Config: config,
		Manager: controllerManager,
	})

	if err != nil {
		logrus.Fatal(err)
	}
	if err := appApi.Run(); err != nil {
		logrus.Fatal(err)
	}
}