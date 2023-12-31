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

func ListFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter productmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := productrepo.NewListFoodRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := productbiz.NewListFoodBiz(repo, requester)

		result, err := biz.ListFood(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
