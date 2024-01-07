package suppliermodel

import "backend/common"

type SupplierCreate struct {
	Id    *string `json:"id" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Email string  `json:"email" gorm:"column:email;"`
	Phone string  `json:"phone" gorm:"column:phone;"`
	Debt  int     `json:"debt" gorm:"column:debt"`
}

func (*SupplierCreate) TableName() string {
	return common.TableSupplier
}

func (data *SupplierCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrSupplierIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrSupplierNameEmpty
	}
	if data.Email != "" && !common.ValidateEmail(data.Email) {
		return ErrSupplierEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrSupplierPhoneInvalid
	}
	if common.ValidateNegativeNumberInt(data.Debt) {
		return ErrSupplierInitDebtInvalid
	}
	return nil
}
