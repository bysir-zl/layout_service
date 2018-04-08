package api

import (
	"github.com/gin-gonic/gin"
	"github.com/bysir-zl/layout/model"
)

// 添加元素
func addComponent(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := model.Component{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.CreateComponent(&s)
	if err != nil {
		return err
	}

	return s
}

// 添加元素
func addComponents(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)
	s := []*model.Component{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.CreateComponents(s)
	if err != nil {
		return err
	}

	return s
}

// 更新元素
func updateComponent(ctx *gin.Context) interface{} {
	//uid := GetUid(ctx)

	s := model.Component{}
	err := ctx.Bind(&s)
	if err != nil {
		return model.ErrBadParams.Append(err.Error())
	}
	err = model.UpdateComponent(s.Id, &s)
	if err != nil {
		return err
	}

	return s
}
