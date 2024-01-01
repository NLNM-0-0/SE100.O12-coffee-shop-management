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
	"log"
	"net/http"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get nearest invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param request query invoicemodel.ReqGetNearestInvoice true "request"
// @Response 200 {object} invoicemodel.ResGetNearestInvoice "list invoice"
// @Response 400 {object} common.AppError "error"
// @Router /invoices/nearest [get]
func GetNearestInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request invoicemodel.ReqGetNearestInvoice
		if err := c.ShouldBind(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		log.Println(request)

		store := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := invoicerepo.NewGetNearestInvoiceRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewGetNearestBiz(repo, requester)

		result, err := biz.GetNearestInvoice(c.Request.Context(), request.AmountNeed)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
