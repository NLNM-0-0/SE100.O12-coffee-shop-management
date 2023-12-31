package ginimportnote

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	importNotes := router.Group("/importNotes", middleware.RequireAuth(appCtx))
	{
		importNotes.GET("", ListImportNote(appCtx))
		importNotes.POST("", CreateImportNote(appCtx))
		importNotes.POST("/:id", ChangeStatusImportNote(appCtx))
		importNotes.GET("/:id", SeeImportNoteDetail(appCtx))
	}
}
