package ginsupplier

import (
	"backend/common"
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/importnote/importnotestore"
	"backend/module/supplier/supplierbiz"
	"backend/module/supplier/suppliermodel/filter"
	"backend/module/supplier/supplierrepo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeSupplierImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var importSupplierFilter filter.SupplierImportFilter
		if err := c.ShouldBind(&importSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierImportNoteRepo(importNoteStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierImportNoteBiz(repo, requester)

		result, err := biz.SeeSupplierImportNote(
			c.Request.Context(), id, &importSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, importSupplierFilter))
	}
}
