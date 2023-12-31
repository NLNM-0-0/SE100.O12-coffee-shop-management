package gincustomer

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/customer/customerbiz"
	"backend/module/customer/customermodel"
	"backend/module/customer/customerrepo"
	"backend/module/invoice/invoicestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeCustomerInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var filter customermodel.FilterInvoice
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		invoiceStore := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := customerrepo.NewSeeCustomerInvoiceRepo(invoiceStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewSeeCustomerInvoiceBiz(repo, requester)

		result, err := biz.SeeCustomerInvoice(c.Request.Context(), id, &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
