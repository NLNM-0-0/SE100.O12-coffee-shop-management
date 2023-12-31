package ginuser

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/user/userbiz"
	"backend/module/user/userrepo"
	"backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeUserDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userStore := userstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := userrepo.NewSeeUserDetailRepo(userStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := userbiz.NewSeeUserDetailBiz(repo, requester)

		result, err := biz.SeeUserDetail(
			c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
