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

func GetAllTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := productrepo.NewGetAllToppingRepo(store)

		biz := productbiz.NewGetAllToppingBiz(repo)

		result, err := biz.GetAllTopping(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
