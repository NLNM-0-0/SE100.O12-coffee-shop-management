package gincustomer

import (
	"backend/common"
	"backend/component/appctx"
	"backend/module/customer/customerbiz"
	"backend/module/customer/customerrepo"
	"backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := customerrepo.NewGetAllCustomerRepo(store)

		biz := customerbiz.NewGetAllCustomerBiz(repo)

		result, err := biz.GetAllCustomer(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
