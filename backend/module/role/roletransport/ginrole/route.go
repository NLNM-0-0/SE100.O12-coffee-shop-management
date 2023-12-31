package ginrole

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	roles := router.Group("/roles", middleware.RequireAuth(appCtx))
	{
		roles.GET("", ListRole(appCtx))
		roles.POST("", CreateRole(appCtx))
		roles.GET("/:id", SeeDetailRole(appCtx))
		roles.PATCH("/:id", UpdateRole(appCtx))
	}
}
