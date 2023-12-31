package ginshopgeneral

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/shopgeneral/shopgeneralbiz"
	"backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewSeeShopGeneralBiz(
			store,
			requester,
		)

		general, err := business.SeeShopGeneral(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(general))
	}
}
