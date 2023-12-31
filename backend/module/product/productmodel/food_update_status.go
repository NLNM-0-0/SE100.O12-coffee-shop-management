package productmodel

import (
	"backend/common"
)

type FoodUpdateStatus struct {
	*ProductUpdateStatus `json:",inline"`
}

func (*FoodUpdateStatus) TableName() string {
	return common.TableFood
}

func (data *FoodUpdateStatus) Validate() error {
	if err := (*data.ProductUpdateStatus).Validate(); err != nil {
		return err
	}
	return nil
}
