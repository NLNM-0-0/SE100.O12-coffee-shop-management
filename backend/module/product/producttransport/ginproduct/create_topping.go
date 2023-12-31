package ginproduct

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/product/productbiz"
	"backend/module/product/productmodel"
	"backend/module/product/productrepo"
	"backend/module/product/productstore"
	"backend/module/recipe/recipestore"
	"backend/module/recipedetail/recipedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.ToppingCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		toppingStore := productstore.NewSQLStore(db)
		recipeStore := recipestore.NewSQLStore(db)
		recipeDetailStore := recipedetailstore.NewSQLStore(db)

		repo := productrepo.NewCreateToppingRepo(
			toppingStore,
			recipeStore,
			recipeDetailStore,
		)

		gen := generator.NewShortIdGenerator()

		business := productbiz.NewCreateToppingBiz(gen, repo, requester)

		if err := business.CreateTopping(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
