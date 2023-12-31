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

func ChangeRoleUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data usermodel.UserUpdateRole

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		userStore := userstore.NewSQLStore(db)
		repo := userrepo.NewChangeRoleUserRepo(userStore)

		business := userbiz.NewChangeRoleUserBiz(repo, requester)

		if err := business.ChangeRoleUser(c.Request.Context(), id, &data); err != nil {
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
