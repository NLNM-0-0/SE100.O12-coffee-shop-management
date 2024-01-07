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

func UpdateInfoIngredient(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data ingredientmodel.IngredientUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := ingredientstore.NewSQLStore(db)

		business := ingredientbiz.NewUpdateInfoIngredientBiz(store, requester)

		if err := business.UpdateInfoIngredient(c.Request.Context(), id, &data); err != nil {
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
