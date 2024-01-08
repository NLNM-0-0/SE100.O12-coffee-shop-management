package main

import (
	"backend/component/appctx"
	"backend/middleware"
	"backend/module/category/categorytransport/gincategory"
	"backend/module/customer/customertransport/gincustomer"
	"backend/module/dashboard/dashboardtransport/gindashboard"
	"backend/module/exportnote/exportnotetransport/ginexportnote"
	"backend/module/feature/featuretransport/ginfeature"
	"backend/module/importnote/importnotetransport/ginimportnote"
	"backend/module/ingredient/ingredienttransport/giningredient"
	"backend/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"backend/module/invoice/invoicetransport/gininvoice"
	"backend/module/product/producttransport/ginproduct"
	"backend/module/role/roletransport/ginrole"
	"backend/module/salereport/salereporttransport/ginsalereport"
	"backend/module/shopgeneral/shopgeneraltransport/ginshopgeneral"
	ginstockreports "backend/module/stockreport/stockreporttransport/ginstockreport"
	"backend/module/supplier/suppliertransport/ginsupplier"
	"backend/module/supplierdebtreport/supplierdebtreporttransport/ginsupplierdebtreport"
	"backend/module/unittype/unittypetransport/ginunittype"
	"backend/module/upload/uploadtransport/ginupload"
	"backend/module/user/usertransport/ginuser"
	cloud "cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type appConfig struct {
	Port      string
	Env       string
	Bucket    string
	DBConnStr string
	SecretKey string
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	fmt.Println("Connecting to database...")
	db, err := connectDatabaseWithRetryIn60s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceAccountKey.json")

	storage, errGetStorage := cloud.NewClient(ctx, sa)
	if errGetStorage != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db, storage, cfg.SecretKey, cfg.Bucket)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	{
		gincategory.SetupRoutes(v1, appCtx)
		gincustomer.SetupRoutes(v1, appCtx)
		ginexportnote.SetupRoutes(v1, appCtx)
		ginfeature.SetupRoutes(v1, appCtx)
		gininvoice.SetupRoutes(v1, appCtx)
		ginimportnote.SetupRoutes(v1, appCtx)
		giningredient.SetupRoutes(v1, appCtx)
		gininventorychecknote.SetupRoutes(v1, appCtx)
		ginproduct.SetupRoutes(v1, appCtx)
		ginrole.SetupRoutes(v1, appCtx)
		ginshopgeneral.SetupRoutes(v1, appCtx)
		ginsupplier.SetupRoutes(v1, appCtx)
		ginuser.SetupRoutes(v1, appCtx)
		ginupload.SetupRoutes(v1, appCtx)
		ginunittype.SetupRoutes(v1, appCtx)
		gindashboard.SetupRoutes(v1, appCtx)
		report := v1.Group("/reports")
		{
			ginstockreports.SetupRoutes(report, appCtx)
			ginsupplierdebtreport.SetupRoutes(report, appCtx)
			ginsalereport.SetupRoutes(report, appCtx)
		}
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalln("Error running server:", err)
	}
}

func loadConfig() (*appConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	return &appConfig{
		Port:      env["PORT"],
		Env:       env["ENVIRONMENT"],
		DBConnStr: env["DB_CONNECTION_STR"],
		SecretKey: env["SECRET_KEY"],
		Bucket:    env["BUCKET"],
	}, nil
}

func connectDatabaseWithRetryIn60s(cfg *appConfig) (*gorm.DB, error) {
	const timeRetry = 60 * time.Second
	var connectDatabase = func(cfg *appConfig) (*gorm.DB, error) {
		db, err := gorm.Open(mysql.Open(cfg.DBConnStr), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 10 seconds: %w", err)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
