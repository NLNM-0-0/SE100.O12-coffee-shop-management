package ginfeature

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	features := router.Group("/features", middleware.RequireAuth(appCtx))
	{
		features.GET("", ListFeature(appCtx))
	}
}
