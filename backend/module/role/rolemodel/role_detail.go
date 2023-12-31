package rolemodel

import "backend/module/feature/featuremodel"

type RoleDetail struct {
	Id       string                       `json:"id" gorm:"column:id;" example:"role id"`
	Name     string                       `json:"name" gorm:"column:name;" example:"admin"`
	Features []featuremodel.FeatureDetail `json:"data" gorm:"-"`
}
