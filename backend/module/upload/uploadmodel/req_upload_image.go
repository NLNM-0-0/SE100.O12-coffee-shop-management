package uploadmodel

import (
	"backend/common"
	"errors"
)

type ReqUploadImage struct {
	ImageType   ImageType `json:"imageType"`
	FolderGroup string    `json:"folderGroup"`
}

var (
	ErrUploadImageInvalid = common.NewCustomError(
		errors.New("image is invalid"),
		"Đinh dạng hình ảnh không đúng",
		"ErrUploadImageInvalid",
	)
)
