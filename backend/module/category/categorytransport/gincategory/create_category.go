package gincategory

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/category/categorybiz"
	"backend/module/category/categorymodel"
	"backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := categorybiz.NewCreateCategoryBiz(gen, store, requester)

		if err := business.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
