package ginunittype

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/unittype/unittypebiz"
	"backend/module/unittype/unittypestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUnitType(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := unittypestore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := unittypebiz.NewGetAllUnitTypeBiz(store)

		result, err := biz.GetAllUnitType(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
