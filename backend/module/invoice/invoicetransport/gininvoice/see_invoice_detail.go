package gininvoice

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/invoice/invoicebiz"
	"backend/module/invoice/invoicerepo"
	"backend/module/invoice/invoicestore"
	"backend/module/invoicedetail/invoicedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeInvoiceDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		invoiceDetailStore := invoicedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		invoiceStore := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := invoicerepo.NewSeeInvoiceDetailRepo(invoiceDetailStore, invoiceStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewSeeInvoiceDetailBiz(
			repo, requester)

		result, err := biz.SeeInvoiceDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
