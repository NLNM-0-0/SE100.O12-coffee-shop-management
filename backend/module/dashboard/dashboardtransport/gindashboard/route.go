package gindashboard

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	stock := router.Group("/dashboard", middleware.RequireAuth(appCtx))
	{
		stock.POST("", SeeDashboard(appCtx))
	}
}
