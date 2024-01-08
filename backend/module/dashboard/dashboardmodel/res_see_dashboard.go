package dashboardmodel

import (
	"backend/module/product/productmodel"
	"time"
)

type ResSeeDashboard struct {
	TimeFrom                     time.Time                     `json:"timeFrom" gorm:"-"`
	TimeTo                       time.Time                     `json:"timeTo" gorm:"-"`
	TotalSale                    int                           `json:"totalSale" gorm:"-"`
	TotalProduct                 int                           `json:"totalProduct" gorm:"-"`
	TotalSold                    int                           `json:"totalSold" gorm:"-"`
	TotalPoint                   int                           `json:"totalPoint" gorm:"-"`
	TopSoldFoods                 []productmodel.FoodWithAmount `json:"topSoldFoods" gorm:"-"`
	ChartSaleComponents          []ChartComponent              `json:"chartSaleComponents" gorm:"-"`
	ChartAmountReceiveComponents []ChartComponent              `json:"chartAmountReceiveComponents" gorm:"-"`
}

type ChartComponent struct {
	Time  time.Time `json:"time" gorm:"-"`
	Value int       `json:"value" gorm:"-"`
}
