package ginproduct

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	foods := router.Group("/foods", middleware.RequireAuth(appCtx))
	{
		foods.GET("", ListFood(appCtx))
		foods.GET("/all", GetAllFood(appCtx))
		foods.POST("", CreateFood(appCtx))
		foods.GET("/:id", SeeDetailFood(appCtx))
		foods.PATCH("/:id", UpdateFood(appCtx))
		foods.PATCH("/status", ChangeStatusFoods(appCtx))
	}
	toppings := router.Group("/toppings", middleware.RequireAuth(appCtx))
	{
		toppings.GET("", ListTopping(appCtx))
		toppings.GET("/all", GetAllTopping(appCtx))
		toppings.POST("", CreateTopping(appCtx))
		toppings.GET("/:id", SeeDetailTopping(appCtx))
		toppings.PATCH("/:id", UpdateTopping(appCtx))
		toppings.PATCH("/status", ChangeStatusToppings(appCtx))
	}
}
