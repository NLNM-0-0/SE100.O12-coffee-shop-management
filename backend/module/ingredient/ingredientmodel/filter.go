package ingredientmodel

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search"`
	MinPrice  *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice  *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	MinAmount *int     `json:"minAmount,omitempty" form:"minAmount"`
	MaxAmount *int     `json:"maxAmount,omitempty" form:"maxAmount"`
	UnitType  string   `json:"unitType,omitempty" form:"unitType"`
}
