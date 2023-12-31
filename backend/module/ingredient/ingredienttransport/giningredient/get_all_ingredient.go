package giningredient

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/ingredient/ingredientbiz"
	"backend/module/ingredient/ingredientstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllIngredient(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := ingredientbiz.NewGetAllIngredientBiz(store)

		result, err := biz.GetAllIngredient(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
