package gincustomer

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/customer/customerbiz"
	"backend/module/customer/customermodel"
	"backend/module/customer/customerrepo"
	"backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data customermodel.CustomerCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := customerstore.NewSQLStore(db)
		repo := customerrepo.NewCreateCustomerRepo(store)

		gen := generator.NewShortIdGenerator()

		business := customerbiz.NewCreateCustomerBiz(gen, repo, requester)

		if err := business.CreateCustomer(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
