package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/bysir-zl/bygo/util/encoder"
	"time"
)

func upload(ctx *gin.Context) interface{} {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	fPath := "static/update/" + encoder.Md5String(file.Filename+time.Now().Format("2006-01-02 15:04:05.999999999"))
	if err := ctx.SaveUploadedFile(file, fPath); err != nil {
		return err
	}
	return fPath
}

func init() {
	// 创建update文件夹
	os.MkdirAll("static/update/", os.ModeDir)
}
