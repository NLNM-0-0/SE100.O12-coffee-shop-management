package ginexportnote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/exportnote/exportnotebiz"
	"backend/module/exportnote/exportnoterepo"
	"backend/module/exportnote/exportnotestore"
	"backend/module/exportnotedetail/exportnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeExportNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		exportNoteDetailStore := exportnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		exportNoteStore := exportnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := exportnoterepo.NewSeeExportNoteDetailRepo(exportNoteDetailStore, exportNoteStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := exportnotebiz.NewSeeExportNoteDetailBiz(
			repo, requester)

		result, err := biz.SeeExportNoteDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
