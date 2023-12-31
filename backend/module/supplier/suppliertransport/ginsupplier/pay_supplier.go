package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/supplierrepo"
	"backend/module/supplier/supplierstore"
	"backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaySupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.SupplierUpdateDebt

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		repo := supplierrepo.NewPaySupplierRepo(supplierStore, supplierDebtStore)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewUpdatePayBiz(gen, repo, requester)

		idSupplierDebt, err := business.PaySupplier(c.Request.Context(), id, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(idSupplierDebt))
	}
}
