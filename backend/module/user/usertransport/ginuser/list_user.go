package ginuser

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/user/userbiz"
	"backend/module/user/usermodel"
	"backend/module/user/userrepo"
	"backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewListUserRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := userbiz.NewListUserBiz(repo, requester)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
