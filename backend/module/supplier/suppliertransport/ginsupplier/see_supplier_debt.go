package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/suppliermodel/filter"
	"backend/module/supplier/supplierrepo"
	"backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeSupplierDebt(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var debtSupplierFilter filter.SupplierDebtFilter
		if err := c.ShouldBind(&debtSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		supplierDebtStore := supplierdebtstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierDebtRepo(supplierDebtStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierDebtBiz(repo, requester)

		result, err := biz.SeeSupplierDebt(
			c.Request.Context(), id, &debtSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, debtSupplierFilter))
	}
}
