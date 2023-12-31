package ginuser

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/component/hasher"
	"backend/middleware"
	"backend/module/user/userbiz"
	"backend/module/user/usermodel"
	"backend/module/user/userrepo"
	"backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		userStore := userstore.NewSQLStore(db)
		repo := userrepo.NewCreateUserRepo(userStore)

		md5 := hasher.NewMd5Hash()
		gen := generator.NewShortIdGenerator()
		biz := userbiz.NewCreateUserBiz(gen, repo, md5, requester)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
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
