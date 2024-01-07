package gininvoice

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/customer/customerstore"
	"backend/module/ingredient/ingredientstore"
	"backend/module/invoice/invoicebiz"
	"backend/module/invoice/invoicemodel"
	"backend/module/invoice/invoicerepo"
	"backend/module/invoice/invoicestore"
	"backend/module/invoicedetail/invoicedetailstore"
	"backend/module/product/productstore"
	"backend/module/shopgeneral/shopgeneralstore"
	"backend/module/sizefood/sizefoodstore"
	"backend/module/stockchangehistory/stockchangehistorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data invoicemodel.InvoiceCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		customerStore := customerstore.NewSQLStore(db)
		sizeFoodStore := sizefoodstore.NewSQLStore(db)
		foodStore := productstore.NewSQLStore(db)
		toppingStore := productstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		shopGeneralStore := shopgeneralstore.NewSQLStore(db)
		stockChangeHistory := stockchangehistorystore.NewSQLStore(db)

		repo := invoicerepo.NewCreateInvoiceRepo(
			invoiceStore,
			invoiceDetailStore,
			customerStore,
			sizeFoodStore,
			foodStore,
			toppingStore,
			ingredientStore,
			shopGeneralStore,
			stockChangeHistory,
		)

		gen := generator.NewShortIdGenerator()

		business := invoicebiz.NewCreateInvoiceBiz(gen, repo, requester)

		if err := business.CreateInvoice(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		shopStore := shopgeneralstore.NewSQLStore(appCtx.GetMainDBConnection())

		printRepo := invoicerepo.NewSeeInvoiceDetailRepo(invoiceDetailStore, invoiceStore)

		biz := invoicebiz.NewPrintInvoiceBiz(
			printRepo, shopStore, requester)

		result, err := biz.PrintInvoice(c.Request.Context(), data.Id)

		if err != nil {
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
