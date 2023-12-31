package ginshopgeneral

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	shopGenerals := router.Group("/shop", middleware.RequireAuth(appCtx))
	{
		shopGenerals.GET("", SeeShopGeneral(appCtx))
		shopGenerals.PATCH("", UpdateShopGeneral(appCtx))
	}
}
