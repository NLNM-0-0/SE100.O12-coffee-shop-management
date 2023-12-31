package ginfeature

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/feature/featurebiz"
	"backend/module/feature/featurestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFeature(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := featurestore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := featurebiz.NewListFeatureBiz(store, requester)

		result, err := biz.ListFeature(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
