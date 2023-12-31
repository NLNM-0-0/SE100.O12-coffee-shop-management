package gincustomer

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/customer/customerbiz"
	"backend/module/customer/customermodel"
	"backend/module/customer/customerrepo"
	"backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateInfoCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data customermodel.CustomerUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := customerstore.NewSQLStore(db)
		repo := customerrepo.NewUpdateInfoCustomerRepo(store)
		biz := customerbiz.NewUpdateInfoCustomerBiz(repo, requester)

		if err := biz.UpdateInfoCustomer(c.Request.Context(), id, &data); err != nil {
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
