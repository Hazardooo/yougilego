package yougilego

type ListResponse[T any] struct {
	Paging struct {
		Limit  int  `json:"limit"`
		Offset int  `json:"offset"`
		Next   bool `json:"next"`
		Count  int  `json:"count"`
	} `json:"paging"`
	Content []T `json:"content"`
}

type IDResponse struct {
	Id string `json:"id"`
}
