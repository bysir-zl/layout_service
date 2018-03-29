package model

import (
	"github.com/go-xorm/xorm"
	"github.com/bysir-zl/layout/config"
)

func SyncAll() (err error) {
	engine, err := xorm.NewEngine("mysql", config.DataSource)
	if err != nil {
		return
	}
	err = engine.Sync2(new(User))
	if err != nil {
		return
	}
	return
}
