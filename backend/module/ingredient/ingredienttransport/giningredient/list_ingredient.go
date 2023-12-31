package giningredient

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/ingredient/ingredientbiz"
	"backend/module/ingredient/ingredientmodel"
	"backend/module/ingredient/ingredientstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListIngredient(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter ingredientmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := ingredientbiz.NewListIngredientBiz(store, requester)

		result, err := biz.ListIngredient(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
