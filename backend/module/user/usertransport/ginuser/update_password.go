package ginuser

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/hasher"
	"backend/middleware"
	"backend/module/user/userbiz"
	"backend/module/user/usermodel"
	"backend/module/user/userrepo"
	"backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		id := requester.GetUserId()

		var data usermodel.UserUpdatePassword

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewUpdatePasswordRepo(store)

		md5 := hasher.NewMd5Hash()
		business := userbiz.NewUpdatePasswordBiz(repo, md5)

		if err := business.UpdatePassword(c.Request.Context(), id, &data); err != nil {
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
