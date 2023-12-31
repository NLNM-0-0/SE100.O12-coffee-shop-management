package giningredient

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	ingredients := router.Group("/ingredients", middleware.RequireAuth(appCtx))
	{
		ingredients.GET("", ListIngredient(appCtx))
		ingredients.GET("/all", GetAllIngredient(appCtx))
		ingredients.POST("", CreateIngredient(appCtx))
	}
}
