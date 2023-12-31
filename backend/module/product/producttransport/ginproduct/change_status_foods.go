package ginproduct

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/product/productbiz"
	"backend/module/product/productmodel"
	"backend/module/product/productrepo"
	"backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusFoods(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.ProductsUpdateStatus

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := productstore.NewSQLStore(db)
		repo := productrepo.NewChangeStatusFoodsRepo(store)
		business := productbiz.NewChangeStatusFoodsBiz(repo, requester)

		if err := business.ChangeStatusFoods(c.Request.Context(), data); err != nil {
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
