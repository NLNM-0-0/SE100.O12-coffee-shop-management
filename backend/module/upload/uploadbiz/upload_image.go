package uploadbiz

import (
	"backend/component/generator"
	"backend/module/upload/uploadmodel"
	cloud "cloud.google.com/go/storage"
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
)

type uploadImageBiz struct {
	gen     generator.IdGenerator
	storage *cloud.Client
}

func NewUploadImageBiz(
	gen generator.IdGenerator,
	storage *cloud.Client) *uploadImageBiz {
	return &uploadImageBiz{
		gen:     gen,
		storage: storage,
	}
}

func (biz *uploadImageBiz) UploadImage(
	ctx context.Context,
	imageType uploadmodel.ImageType,
	bucket string,
	fileHeader *multipart.FileHeader) (*string, error) {
	fileNameLocal := fileHeader.Filename
	extension := filepath.Ext(fileNameLocal)
	if !uploadmodel.CheckImageExtension(extension) {
		return nil, uploadmodel.ErrUploadImageInvalid
	}

	id, err := biz.gen.GenerateId()
	if err != nil {
		return nil, err
	}

	file, errOpenFile := fileHeader.Open()
	if errOpenFile != nil {
		return nil, err
	}
	defer file.Close()

	fileName := filepath.Base(fileNameLocal[:len(fileNameLocal)-len(filepath.Ext(fileNameLocal))]) + id
	filePath := string(imageType) + "/" + fileName
	imageUrl := "https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/" + string(imageType) + "%2F" + fileName + "?alt=media"

	wc := biz.storage.Bucket(bucket).Object(filePath).NewWriter(ctx)
	_, err = io.Copy(wc, file)
	if err != nil {
		return nil, err
	}

	if err := wc.Close(); err != nil {
		return nil, err
	}

	return &imageUrl, nil
}
