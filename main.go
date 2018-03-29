package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bysir-zl/layout/api"
	"github.com/bysir-zl/layout/config"
	"github.com/bysir-zl/layout/model"
	_ "github.com/bysir-zl/layout/worker"
	"time"
	"log"
	"os"
)

func main() {
	ParseTool()

	Init()

	runMode := gin.DebugMode
	if !config.Debug {
		runMode = gin.ReleaseMode
	}
	gin.SetMode(runMode)

	router := gin.Default()
	// 跨域
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Token,Content-Type")
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS,PUT,POST,GET,DELETE")
	})
	router.NoRoute(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(200, map[string]string{"status": "ok"})
		}
	})
	router.Static("/static", "./static")
	api.RegisterAll(router)

	router.Run(config.Listen)
}

func ParseTool() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "sync":
			err := model.SyncAll()
			if err != nil {
				log.Print("SyncAll error:", err)
				break
			}
			log.Print("SyncAll success")
		}
		os.Exit(0)
	}
}

// 初始化app
func Init() {
	time.AfterFunc(1*time.Second, func() {
		//_, err := model.GetApp()
		//if err != nil {
		//	if !model.ErrNotFind.Is(err) {
		//		log.Panicf("获取app信息失败 %v", err)
		//		return
		//	}
		//
		//	model.CreateApp(&model.App{
		//		Id:                    1,
		//		DonationPoint:         0,
		//		PointRatioForMoney:    1,
		//		WeightMinAppointment:  30000,
		//		CompleteUserInfoPoint: 5,
		//		ServicePhone:          "",
		//	})
		//}

		// 如果没找到管理员1(root)就创建一个
		//_, err := model.GetAdmin(1)
		//if model.ErrNotFind.Is(err) {
		//	a := model.Administrator{
		//		Point:    9999,
		//		Id:       1,
		//		Nickname: "root",
		//		Phone:    "15288888888",
		//		Pwd:      "123456",
		//		Type:     model.AdministratorTypeSuper,
		//		CanLogin: true,
		//		Sex:      1,
		//	}
		//	_, err := model.CreateAdmin(&a)
		//	if err != nil {
		//		log.Panicf("初始化root管理员失败 %v", err)
		//		return
		//	}
		//}

	})
}
