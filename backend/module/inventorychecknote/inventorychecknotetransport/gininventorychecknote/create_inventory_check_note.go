package gininventorychecknote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/ingredient/ingredientstore"
	"backend/module/inventorychecknote/inventorychecknotebiz"
	"backend/module/inventorychecknote/inventorychecknotemodel"
	"backend/module/inventorychecknote/inventorychecknoterepo"
	"backend/module/inventorychecknote/inventorychecknotestore"
	"backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data inventorychecknotemodel.InventoryCheckNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		inventoryCheckNoteStore := inventorychecknotestore.NewSQLStore(db)
		inventoryCheckNoteDetailStore := inventorychecknotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)

		repo := inventorychecknoterepo.NewCreateInventoryCheckNoteRepo(
			inventoryCheckNoteStore,
			inventoryCheckNoteDetailStore,
			ingredientStore,
		)

		gen := generator.NewShortIdGenerator()

		business := inventorychecknotebiz.NewCreateInventoryCheckNoteBiz(gen, repo, requester)

		if err := business.CreateInventoryCheckNote(c.Request.Context(), &data); err != nil {
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
