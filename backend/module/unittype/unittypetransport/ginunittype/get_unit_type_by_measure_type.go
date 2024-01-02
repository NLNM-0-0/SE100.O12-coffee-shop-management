package ginunittype

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/unittype/unittypebiz"
	"backend/module/unittype/unittypemodel"
	"backend/module/unittype/unittypestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUnitTypeByMeasureType(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		measureType := c.Param("measureType")
		if measureType != string(unittypemodel.Weight) &&
			measureType != string(unittypemodel.Volume) &&
			measureType != string(unittypemodel.Unit) {
			panic(common.ErrInvalidRequest(unittypemodel.ErrUnitTypeMeasureTypeInvalid))
		}

		store := unittypestore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := unittypebiz.NewGetUnitTypeByMeasureTypeBiz(store)

		result, err := biz.GetUnitTypeByMeasureType(c.Request.Context(), unittypemodel.MeasureType(measureType))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
