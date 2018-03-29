package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/go-xorm/core"
	"github.com/bysir-zl/layout/config"
)

var engine *xorm.Engine

func init() {
	var err error
	dataSourceName := config.DataSource
	if dataSourceName == "" {
		panic("dataSource is undefined in config")
	}
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	engine.SetMapper(core.SnakeMapper{})
	engine.ShowSQL(config.Debug)
}
