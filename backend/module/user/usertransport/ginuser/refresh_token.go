package ginuser

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/tokenprovider/jwt"
	"backend/module/user/userbiz"
	"backend/module/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRefreshToken
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		business := userbiz.NewRefreshTokenBiz(60*60*24, tokenProvider)
		account, err := business.RefreshToken(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
