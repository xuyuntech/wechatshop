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
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name: "config-file",
		Value: "./config.yml",
		EnvVar: "CONFIG_FILE",
	},
}


func main(){
	app := cli.NewApp()
	app.Name = "xuyuntech eShop"
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func action(c *cli.Context) error {

	configFilePath := c.String("config-file")
	if len(configFilePath) == 0 {
		return errors.New("no config file provided")
	}

	yConfig, err := config.LoadConfig(configFilePath)
	if err != nil {
		return err
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
		return fmt.Errorf("new db engine err: %v", err)
	}

	controllerManager, err := manager.NewDefaultManager(yConfig, engine)
	if err != nil {
		return fmt.Errorf("new controller manager err: %v", err)
	}

	appApi, err := api.NewApi(&api.ApiConfig{
		Config: yConfig,
		Manager: controllerManager,
	})

	if err != nil {
		return err
	}
	if err := appApi.Run(); err != nil {
		return err
	}

	return nil
}