package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/supplierrepo"
	"backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeSupplierDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		supplierStore := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierDetailRepo(supplierStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierDetailBiz(repo, requester)

		result, err := biz.SeeSupplierDetail(
			c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
