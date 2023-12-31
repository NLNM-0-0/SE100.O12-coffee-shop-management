package ginrole

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/role/rolebiz"
	"backend/module/role/rolerepo"
	"backend/module/role/rolestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := rolestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := rolerepo.NewListRoleRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewListRoleBiz(repo, requester)

		result, err := biz.ListRole(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
