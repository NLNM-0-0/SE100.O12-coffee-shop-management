package ginstockreports

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/ingredient/ingredientstore"
	"backend/module/stockchangehistory/stockchangehistorystore"
	"backend/module/stockreport/stockreportbiz"
	"backend/module/stockreport/stockreportmodel"
	"backend/module/stockreport/stockreportstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindStockReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data stockreportmodel.ReqFindStockReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		ingredientStore := ingredientstore.NewSQLStore(db)
		stockChangeHistoryStore := stockchangehistorystore.NewSQLStore(db)
		stockReportStore := stockreportstore.NewSQLStore(db)

		gen := generator.NewShortIdGenerator()
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := stockreportbiz.NewFindStockReportBiz(
			gen,
			ingredientStore,
			stockChangeHistoryStore,
			stockReportStore,
			requester,
		)

		report, err := business.FindStockReport(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
