package invoicemodel

import "backend/module/shopgeneral/shopgeneralmodel"

type InvoicePrint struct {
	Invoice     Invoice                      `json:"invoice" gorm:"foreignkey:invoiceId;association_foreignkey:id"`
	ShopGeneral shopgeneralmodel.ShopGeneral `json:"shop" gorm:"foreignkey:invoiceId;association_foreignkey:id"`
}
