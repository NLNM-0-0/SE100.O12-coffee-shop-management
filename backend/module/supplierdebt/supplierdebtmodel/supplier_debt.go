package supplierdebtmodel

import (
	"backend/common"
	"backend/common/enum"
	"backend/module/user/usermodel"
	"time"
)

type SupplierDebt struct {
	Id            string               `json:"id" gorm:"column:id;"`
	SupplierId    string               `json:"supplierId" gorm:"column:supplierId;"`
	Amount        int                  `json:"amount" gorm:"column:amount;"`
	AmountLeft    int                  `json:"amountLeft" gorm:"column:amountLeft;"`
	DebtType      *enum.DebtType       `json:"type" gorm:"column:type;"`
	CreatedBy     string               `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt     *time.Time           `json:"createdAt" gorm:"column:createdAt;"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}
