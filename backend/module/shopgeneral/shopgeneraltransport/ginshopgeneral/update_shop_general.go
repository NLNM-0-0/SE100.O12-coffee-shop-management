package ginshopgeneral

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/shopgeneral/shopgeneralbiz"
	"backend/module/shopgeneral/shopgeneralmodel"
	"backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data shopgeneralmodel.ShopGeneralUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewUpdateGeneralShopBiz(
			store,
			requester,
		)

		if err := business.UpdateGeneralShop(c.Request.Context(), &data); err != nil {
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
