package dashboardmodel

import (
	"backend/module/product/productmodel"
	"time"
)

type ResSeeDashboard struct {
	TimeFrom             time.Time                     `json:"timeFrom" gorm:"-"`
	TimeTo               time.Time                     `json:"timeTo" gorm:"-"`
	TotalSale            int                           `json:"totalSale" gorm:"-"`
	TotalSold            int                           `json:"totalSold" gorm:"-"`
	TotalCustomer        int                           `json:"totalCustomer" gorm:"-"`
	TotalPoint           int                           `json:"totalPoint" gorm:"-"`
	TopSoldFoods         []productmodel.FoodWithAmount `json:"topSoldFoods" gorm:"-"`
	ChartPriceComponents []ChartComponent              `json:"chartPriceComponents" gorm:"-"`
	ChartCostComponents  []ChartComponent              `json:"chartCostComponents" gorm:"-"`
}

type ChartComponent struct {
	Time  time.Time `json:"time" gorm:"-"`
	Value int       `json:"value" gorm:"-"`
}
