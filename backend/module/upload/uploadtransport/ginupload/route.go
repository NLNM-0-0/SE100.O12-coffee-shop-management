package ginupload

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	uploadFile := router.Group("/upload", middleware.RequireAuth(appCtx))
	uploadFile.POST("", UploadFile(appCtx))
}
