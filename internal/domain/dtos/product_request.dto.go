package dtos

type ProductRequestDto struct {
	SearchQuery string `json:"search_query"`
	CategoryIDs []int  `json:"category_ids"`
}
