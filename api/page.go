package api

import (
	"github.com/gin-gonic/gin"
	"github.com/bysir-zl/layout/model"
)

// 添加元素
func updatePage(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := model.Page{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.UpdatePage(s.Id, &s)
	if err != nil {
		return err
	}

	return "ok"
}

// 添加元素
func getPageLayout(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := IdParams{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	l,err := model.GetLayoutPage(s.Id)
	if err != nil {
		return err
	}

	return l
}
