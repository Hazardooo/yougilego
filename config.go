package yougilego

type YGConfig struct {
	BugTruckerProjectName string `json:"bugTruckerProjectName"`
	BugTruckerBoardName   string `json:"bugTruckerBoardName"`
	BugTruckerColumnName  string `json:"bugTruckerColumnName"`
}

type SuccessResponse struct {
	Id string `json:"id"`
}

type ListResponse[T any] struct {
	Paging struct {
		Limit  int  `json:"limit"`
		Offset int  `json:"offset"`
		Next   bool `json:"next"`
		Count  int  `json:"count"`
	} `json:"paging"`
	Content []T `json:"content"`
}
