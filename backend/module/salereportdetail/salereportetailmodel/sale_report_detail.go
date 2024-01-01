package salereportetailmodel

import "backend/module/product/productmodel"

type SaleReportDetail struct {
	FoodId     string                  `json:"-"`
	Food       productmodel.SimpleFood `json:"food"`
	Amount     int                     `json:"amount"`
	TotalSales int                     `json:"totalSales"`
}
