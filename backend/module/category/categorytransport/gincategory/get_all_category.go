package gincategory

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/category/categorybiz"
	"backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := categorybiz.NewGetAllCategoryBiz(store)

		result, err := biz.GetAllCategory(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
