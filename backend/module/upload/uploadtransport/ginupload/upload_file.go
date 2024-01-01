package ginupload

import (
	"backend/common"
	"backend/component/appctx"
	"backend/component/generator"
	"backend/module/upload/uploadbiz"
	"backend/module/upload/uploadmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imageType := c.DefaultPostForm("imageType", string(uploadmodel.Other))
		if imageType != string(uploadmodel.Other) &&
			imageType != string(uploadmodel.Avatar) &&
			imageType != string(uploadmodel.Food) &&
			imageType != string(uploadmodel.Topping) {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		biz := uploadbiz.NewUploadImageBiz(gen, appCtx.GetStorage())

		url, err := biz.UploadImage(
			c.Request.Context(), uploadmodel.ImageType(imageType), appCtx.GetBucket(), fileHeader)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(url))
	}
}
