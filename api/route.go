package api

import "github.com/gin-gonic/gin"

func RegisterAll(e *gin.Engine) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/item", toGinHandler(addComponent, ""))
		//v1.GET("/item", toGinHandler(addItem, ""))
		v1.PUT("/item", toGinHandler(updateComponent, ""))
		v1.POST("/item/multi", toGinHandler(addComponent, ""))

		v1.POST("/update", toGinHandler(upload, ""))
		v1.PUT("/page", toGinHandler(updatePage, ""))
		v1.GET("/page/layout", toGinHandler(getPageLayout, ""))
	}

}
