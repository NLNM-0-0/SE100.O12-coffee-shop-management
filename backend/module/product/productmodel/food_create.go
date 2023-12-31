package productmodel

import (
	"backend/common"
	"backend/module/sizefood/sizefoodmodel"
)

type FoodCreate struct {
	*ProductCreate `json:",inline"`
	Categories     []string                       `json:"categories" gorm:"-"`
	Sizes          []sizefoodmodel.SizeFoodCreate `json:"sizes" gorm:"-"`
	Image          *string                        `json:"image" gorm:"column:image;"`
}

func (*FoodCreate) TableName() string {
	return common.TableFood
}

func (data *FoodCreate) Validate() error {
	if data.ProductCreate == nil {
		return ErrFoodProductInfoEmpty
	}
	if err := (*data.ProductCreate).Validate(); err != nil {
		return err
	}
	if !common.ValidateUrl(data.Image, common.DefaultImageFood) {
		return ErrFoodImageInvalid
	}
	if data.Categories == nil || len(data.Categories) == 0 {
		return ErrFoodCategoryEmpty
	}
	mapExistCategory := make(map[string]int)
	for _, v := range data.Categories {
		mapExistCategory[v]++
		if mapExistCategory[v] == 2 {
			return ErrFoodExistDuplicateCategory
		}
	}
	if data.Sizes == nil || len(data.Sizes) == 0 {
		return ErrFoodSizeEmpty
	}
	for _, v := range data.Sizes {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
