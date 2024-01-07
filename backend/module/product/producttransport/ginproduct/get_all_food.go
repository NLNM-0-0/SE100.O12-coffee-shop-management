package ginproduct

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/product/productbiz"
	"backend/module/product/productrepo"
	"backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := productrepo.NewGetAllFoodRepo(store)

		biz := productbiz.NewGetAllFoodBiz(repo)

		result, err := biz.GetAllFood(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
