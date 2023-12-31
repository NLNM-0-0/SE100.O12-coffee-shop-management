package ginrole

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/role/rolebiz"
	"backend/module/role/rolemodel"
	"backend/module/role/rolerepo"
	"backend/module/role/rolestore"
	"backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data rolemodel.RoleCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)

		repo := rolerepo.NewCreateRoleRepo(
			roleStore,
			roleFeatureStore,
		)

		gen := generator.NewShortIdGenerator()

		business := rolebiz.NewCreateRoleStore(gen, repo, requester)

		if err := business.CreateRole(c.Request.Context(), &data); err != nil {
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
