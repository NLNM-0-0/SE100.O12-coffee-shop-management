package gincustomer

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	customers := router.Group("/customers", middleware.RequireAuth(appCtx))
	{
		customers.GET("", ListCustomer(appCtx))
		customers.POST("", CreateCustomer(appCtx))
		customers.GET("/:id", SeeCustomerDetail(appCtx))
		customers.GET("/:id/invoices", SeeCustomerInvoice(appCtx))
		customers.POST("/:id", UpdateInfoCustomer(appCtx))
	}
}
