package importnotemodel

type Filter struct {
	SearchKey         string   `json:"searchKey,omitempty" form:"search"`
	MinPrice          *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice          *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	DateFromCreatedAt *int64   `json:"createdAtFrom,omitempty" form:"createdAtFrom"`
	DateToCreatedAt   *int64   `json:"createdAtTo,omitempty" form:"createdAtTo"`
	DateFromClosedAt  *int64   `json:"closedAtFrom,omitempty" form:"closedAtFrom"`
	DateToClosedAt    *int64   `json:"closedAtTo,omitempty" form:"closedAtTo"`
	Supplier          *string  `json:"supplier,omitempty" form:"supplier"`
	CreatedBy         *string  `json:"createdBy,omitempty" form:"createdBy"`
	ClosedBy          *string  `json:"closedBy,omitempty" form:"closedBy"`
	Status            string   `json:"status,omitempty" form:"status"`
}
