package ginunittype

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	unitTypes := router.Group("/unitTypes", middleware.RequireAuth(appCtx))
	{
		unitTypes.GET("", GetAllUnitType(appCtx))
		unitTypes.GET("/:measureType", GetUnitTypeByMeasureType(appCtx))
	}
}
