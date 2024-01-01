package appctx

import (
	cloud "cloud.google.com/go/storage"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetStorage() *cloud.Client
	GetSecretKey() string
	GetBucket() string
}

type appCtx struct {
	db        *gorm.DB
	storage   *cloud.Client
	secretKey string
	bucket    string
}

func NewAppContext(
	db *gorm.DB,
	storage *cloud.Client,
	secretKey string,
	bucket string) *appCtx {
	return &appCtx{
		db:        db,
		storage:   storage,
		secretKey: secretKey,
		bucket:    bucket,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetStorage() *cloud.Client {
	return ctx.storage
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetBucket() string {
	return ctx.bucket
}
