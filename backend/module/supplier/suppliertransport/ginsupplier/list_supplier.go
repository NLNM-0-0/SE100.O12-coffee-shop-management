package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/suppliermodel/filter"
	"backend/module/supplier/supplierrepo"
	"backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filterSupplier filter.Filter
		if err := c.ShouldBind(&filterSupplier); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := supplierrepo.NewListSupplierRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewListSupplierRepo(repo, requester)

		result, err := biz.ListSupplier(c.Request.Context(), &filterSupplier, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filterSupplier))
	}
}
