package ginimportnote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/importnote/importnotebiz"
	"backend/module/importnote/importnotemodel"
	"backend/module/importnote/importnoterepo"
	"backend/module/importnote/importnotestore"
	"backend/module/importnotedetail/importnotedetailstore"
	"backend/module/ingredient/ingredientstore"
	"backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ImportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)

		repo := importnoterepo.NewCreateImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			ingredientStore,
			supplierStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewCreateImportNoteBiz(gen, repo, requester)

		if err := business.CreateImportNote(c.Request.Context(), &data); err != nil {
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
