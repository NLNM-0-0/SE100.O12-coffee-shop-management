package ginstockreports

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	stock := router.Group("/stock", middleware.RequireAuth(appCtx))
	{
		stock.POST("", FindStockReport(appCtx))
	}

}
