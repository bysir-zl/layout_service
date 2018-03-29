package api

import (
	"github.com/gin-gonic/gin"
	"github.com/bysir-zl/layout/model"
)

// 添加元素
func createdItem(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := model.Item{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.CreateItem(&s)
	if err != nil {
		return err
	}

	return s
}

// 添加元素
func updateItem(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := model.Item{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.UpdateItem(s.Id,&s)
	if err != nil {
		return err
	}

	return s
}
