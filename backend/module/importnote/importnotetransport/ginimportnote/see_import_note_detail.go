package ginimportnote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/importnote/importnotebiz"
	"backend/module/importnote/importnoterepo"
	"backend/module/importnote/importnotestore"
	"backend/module/importnotedetail/importnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeImportNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := importnoterepo.NewSeeImportNoteDetailRepo(importNoteStore, importNoteDetailStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeImportNoteDetailBiz(
			repo, requester)

		result, err := biz.SeeImportNoteDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
