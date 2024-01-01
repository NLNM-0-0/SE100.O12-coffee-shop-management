package gininvoice

import (
	"backend/component/appctx"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	invoices := router.Group("/invoices", middleware.RequireAuth(appCtx))
	{
		invoices.GET("", ListInvoice(appCtx))
		invoices.GET("/nearest", GetNearestInvoice(appCtx))
		invoices.POST("", CreateInvoice(appCtx))
		invoices.GET("/:id", SeeInvoiceDetail(appCtx))
	}
}
