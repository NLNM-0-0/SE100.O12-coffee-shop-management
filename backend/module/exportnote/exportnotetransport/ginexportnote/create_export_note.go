package ginexportnote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/exportnote/exportnotebiz"
	"backend/module/exportnote/exportnotemodel"
	"backend/module/exportnote/exportnoterepo"
	"backend/module/exportnote/exportnotestore"
	"backend/module/exportnotedetail/exportnotedetailstore"
	"backend/module/ingredient/ingredientstore"
	"backend/module/stockchangehistory/stockchangehistorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data exportnotemodel.ExportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		exportNoteStore := exportnotestore.NewSQLStore(db)
		exportNoteDetailStore := exportnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)

		repo := exportnoterepo.NewCreateExportNoteRepo(
			exportNoteStore,
			exportNoteDetailStore,
			ingredientStore,
			stockChangeHistoryStore,
		)

		gen := generator.NewShortIdGenerator()

		business := exportnotebiz.NewCreateExportNoteBiz(gen, repo, requester)

		if err := business.CreateExportNote(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
