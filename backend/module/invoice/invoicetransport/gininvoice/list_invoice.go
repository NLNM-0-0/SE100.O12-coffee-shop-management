package gininvoice

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/invoice/invoicebiz"
	"backend/module/invoice/invoicemodel"
	"backend/module/invoice/invoicerepo"
	"backend/module/invoice/invoicestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter invoicemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := invoicerepo.NewListImportNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewListInvoiceBiz(repo, requester)

		result, err := biz.ListInvoice(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
