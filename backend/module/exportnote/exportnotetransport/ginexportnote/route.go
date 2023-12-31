package ginexportnote

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	exportNotes := router.Group("/exportNotes", middleware.RequireAuth(appCtx))
	{
		exportNotes.GET("", ListExportNote(appCtx))
		exportNotes.POST("", CreateExportNote(appCtx))
		exportNotes.GET("/:id", SeeExportNoteDetail(appCtx))
	}
}
