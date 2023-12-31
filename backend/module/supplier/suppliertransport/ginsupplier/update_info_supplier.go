package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/supplierrepo"
	"backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateInfoSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.SupplierUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := supplierstore.NewSQLStore(db)
		repo := supplierrepo.NewUpdateInfoSupplierRepo(store)

		business := supplierbiz.NewUpdateInfoSupplierBiz(repo, requester)

		if err := business.UpdateInfoSupplier(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
