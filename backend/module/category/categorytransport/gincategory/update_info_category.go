package gincategory

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/category/categorybiz"
	"backend/module/category/categorymodel"
	"backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateInfoCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data categorymodel.CategoryUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := categorybiz.NewUpdateInfoCategoryBiz(store, requester)

		if err := biz.UpdateInfoCategory(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
