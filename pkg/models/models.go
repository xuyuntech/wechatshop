package models

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"net/url"
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/go-xorm/core"
	"github.com/robfig/cron"
)

type DBOptions struct {
	User string
	Password string
	Host string
	Port int
	Name string
}

func NewEngine(config DBOptions, t []interface{}) (*xorm.Engine, error){
	var Param string = "?"
	//var _tables []interface{}
	if strings.Contains(config.Name, Param) {
		Param = "&"
	}
	var connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s",
		url.QueryEscape(config.User),
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
		config.Name,"Asia%2FShanghai")

	logrus.Infof("Connect to db: %s", connStr)
	x, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		return nil,err
	}
	logrus.Info("Connect to db ok.")
	x.SetMapper(core.GonicMapper{})
	logrus.Infof("start to sync tables ...")
	if err = x.StoreEngine("InnoDB").Sync2(t...); err != nil {
		return nil, fmt.Errorf("sync tables err: %v",err)
	}
	x.ShowSQL(true)
	go ping(x)
	return x, nil
}


func Tables() []interface{} {
	var tables []interface{}
	tables = append(tables, new(User))
	return tables
}

func ping(engine *xorm.Engine){
	logrus.Debugf("start to pint db engine.")
	forever := make(chan bool)
	c := cron.New()
	c.AddFunc("@every 1m", func(){
		if err := engine.Ping(); err != nil {
			logrus.Errorf("ping err: %s", err.Error())
		}
	})
	c.Start()
	<-forever
}
