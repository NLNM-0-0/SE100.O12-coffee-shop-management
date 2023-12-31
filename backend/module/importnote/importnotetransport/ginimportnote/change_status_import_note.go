package ginimportnote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/importnote/importnotebiz"
	"backend/module/importnote/importnotemodel"
	"backend/module/importnote/importnoterepo"
	"backend/module/importnote/importnotestore"
	"backend/module/importnotedetail/importnotedetailstore"
	"backend/module/ingredient/ingredientstore"
	"backend/module/supplier/supplierstore"
	"backend/module/supplierdebt/supplierdebtstore"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idImportNote := c.Param("id")
		if idImportNote == "" {
			panic(common.ErrInvalidRequest(errors.New("param id not exist")))
		}

		var data importnotemodel.ImportNoteUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.ClosedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)

		repo := importnoterepo.NewChangeStatusImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			ingredientStore,
			supplierStore,
			supplierDebtStore,
		)

		business := importnotebiz.NewChangeStatusImportNoteBiz(repo, requester)

		if err := business.ChangeStatusImportNote(
			c.Request.Context(),
			idImportNote,
			&data,
		); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
