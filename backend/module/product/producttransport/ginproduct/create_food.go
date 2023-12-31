package ginproduct

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/category/categorystore"
	"backend/module/categoryfood/categoryfoodstore"
	"backend/module/product/productbiz"
	"backend/module/product/productmodel"
	"backend/module/product/productrepo"
	"backend/module/product/productstore"
	"backend/module/recipe/recipestore"
	"backend/module/recipedetail/recipedetailstore"
	"backend/module/sizefood/sizefoodstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		foodStore := productstore.NewSQLStore(db)
		categoryFoodStore := categoryfoodstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)
		sizeFoodStore := sizefoodstore.NewSQLStore(db)
		recipeStore := recipestore.NewSQLStore(db)
		recipeDetailStore := recipedetailstore.NewSQLStore(db)

		repo := productrepo.NewCreateFoodRepo(
			foodStore,
			categoryFoodStore,
			categoryStore,
			sizeFoodStore,
			recipeStore,
			recipeDetailStore,
		)

		gen := generator.NewShortIdGenerator()

		business := productbiz.NewCreateFoodBiz(gen, repo, requester)

		if err := business.CreateFood(c.Request.Context(), &data); err != nil {
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
