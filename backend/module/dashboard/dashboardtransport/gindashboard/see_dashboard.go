package gindashboard

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/dashboard/dashboardbiz"
	"backend/module/dashboard/dashboardmodel"
	"backend/module/dashboard/dashboardrepo"
	"backend/module/invoice/invoicestore"
	"backend/module/invoicedetail/invoicedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDashboard(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dashboardmodel.ReqSeeDashboard

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)

		dashboardRepo := dashboardrepo.NewSeeDashboardBiz(invoiceStore, invoiceDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := dashboardbiz.NewSeeDashboardBiz(
			dashboardRepo, requester,
		)

		dashboard, err := business.SeeDashboard(
			c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dashboard))
	}
}
