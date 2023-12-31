package gininventorychecknote

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/inventorychecknote/inventorychecknotebiz"
	"backend/module/inventorychecknote/inventorychecknoterepo"
	"backend/module/inventorychecknote/inventorychecknotestore"
	"backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventoryCheckNoteId := c.Param("id")

		inventoryCheckNoteStore :=
			inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		inventoryCheckNoteDetailStore :=
			inventorychecknotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewSeeDetailInventoryCheckNoteRepo(
			inventoryCheckNoteStore, inventoryCheckNoteDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailInventoryCheckNote(
			c.Request.Context(), inventoryCheckNoteId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
