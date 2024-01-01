package supplierdebtreportdetailmodel

import (
	"backend/common"
	"backend/module/supplier/suppliermodel"
)

type SupplierDebtReportDetail struct {
	ReportId   string                       `json:"-" gorm:"column:reportId;"`
	SupplierId string                       `json:"-" gorm:"column:supplierId;"`
	Supplier   suppliermodel.SimpleSupplier `json:"supplier"`
	Initial    int                          `json:"initial" gorm:"column:initial;" example:"100000"`
	Debt       int                          `json:"debt" gorm:"column:debt;" example:"-40000"`
	Pay        int                          `json:"pay" gorm:"column:pay;" example:"20000"`
	Final      int                          `json:"final" gorm:"column:final;" example:"80000"`
}

func (*SupplierDebtReportDetail) TableName() string {
	return common.TableSupplierDebtReportDetail
}
