package ginproduct

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/product/productbiz"
	"backend/module/product/productrepo"
	"backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := productrepo.NewSeeDetailFoodRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := productbiz.NewSeeDetailFoodBiz(repo, requester)

		result, err := biz.SeeDetailFood(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
