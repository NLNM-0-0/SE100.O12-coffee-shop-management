package ginshopgeneral

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/shopgeneral/shopgeneralbiz"
	"backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewSeeShopGeneralBiz(
			store,
		)

		general, err := business.SeeShopGeneral(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(general))
	}
}
