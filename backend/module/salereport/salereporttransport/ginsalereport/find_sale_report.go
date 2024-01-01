package ginsalereport

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/invoice/invoicestore"
	"backend/module/invoicedetail/invoicedetailstore"
	"backend/module/salereport/salereportbiz"
	"backend/module/salereport/salereportmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindSaleReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data salereportmodel.ReqFindSaleReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := salereportbiz.NewFindSaleReportBiz(
			invoiceStore, invoiceDetailStore, requester,
		)

		report, err := business.FindSaleReport(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
