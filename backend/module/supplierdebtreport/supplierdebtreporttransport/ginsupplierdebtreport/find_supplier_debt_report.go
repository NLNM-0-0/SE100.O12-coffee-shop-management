package ginsupplierdebtreport

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/supplier/supplierstore"
	"backend/module/supplierdebt/supplierdebtstore"
	"backend/module/supplierdebtreport/supplierdebtreportbiz"
	"backend/module/supplierdebtreport/supplierdebtreportmodel"
	"backend/module/supplierdebtreport/supplierdebtreportstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindSupplierDebtReport(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data supplierdebtreportmodel.ReqFindSupplierDebtReport

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		supplierDebtReportStore := supplierdebtreportstore.NewSQLStore(db)

		gen := generator.NewShortIdGenerator()
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := supplierdebtreportbiz.NewFindSupplierDebtReportBiz(
			gen,
			supplierStore,
			supplierDebtStore,
			supplierDebtReportStore,
			requester,
		)

		report, err := business.FindSupplierDebtReport(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(report))
	}
}
