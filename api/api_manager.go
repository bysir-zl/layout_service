package api

import (
	"github.com/gin-gonic/gin"
	"github.com/bysir-zl/bygo/util/auth"
	"strconv"
	"errors"
	"github.com/bysir-zl/layout/config"
	"time"
	"github.com/bysir-zl/bygo/util/encoder"
	"github.com/bysir-zl/layout/model"
)

type Api func(ctx *gin.Context) interface{}

var Auth = map[string]func(ctx *gin.Context) error{
	"user": func(ctx *gin.Context) error {
		token := ctx.GetHeader("Token")
		jwt, errCode := auth.JWTDecode(config.Key, token)
		if errCode != 0 {
			if config.Debug {
				var id int64 = 1
				ctx.Set("uid", id)
				return nil
			}
			return model.ErrNotAuth.Append(auth.ErrcodeString(errCode))
		}
		if jwt.Typ != "user" {
			return model.ErrNotAuth.Append("bad typ:" + jwt.Typ)
		}

		id, _ := strconv.ParseInt(jwt.Sub, 10, 64)
		if id == 0 {
			return model.ErrNotAuth.Append("bad token: " + token)
		}
		u, _ := model.GetUser(id)
		if u == nil || u.Id == 0 {
			return model.ErrNotAuth.Append("user is not exist")
		}

		ctx.Set("uid", id)
		return nil
	},

	"client": func(ctx *gin.Context) error { // 简单的客户端密码认证
		if config.Debug {
			return nil
		}
		pwd := ctx.GetHeader("PWD")
		if pwd == "" {
			pwd, _ = ctx.GetQuery("PWD")
			if pwd == "" {
				return errors.New("PWD is empty")
			}
		}

		now := time.Now().Format("1504") // 时分
		pre := time.Now().Add(-time.Minute).Format("1504")
		next := time.Now().Add(time.Minute).Format("1504")
		for _, t := range []string{now, pre, next} {
			sign := encoder.Md5String(config.Key + t)
			if sign == pwd {
				return nil
			}
		}

		return errors.New("pwd err")
	},
}

type ErrCoder interface {
	Coder() int
	Error() string
}

func toGinHandler(api Api, authMethod string) (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		var rsp interface{}
		if authMethod != "" {
			if authFunc, ok := Auth[authMethod]; ok {
				err := authFunc(ctx)
				if err != nil {
					rsp = err
				}
			}
		}
		if rsp == nil {
			rsp = api(ctx)
			if rsp == nil {
				return
			}
		}

		switch x := rsp.(type) {
		case ErrCoder:
			ctx.JSON(400, map[string]interface{}{
				"code": x.Coder(),
				"msg":  x.Error(),
			})
			return
		case error:
			ctx.JSON(400, map[string]interface{}{
				"code": -1,
				"msg":  x.Error(),
			})
			return
		default:
			ctx.JSON(200, rsp)
		}
	}
}
