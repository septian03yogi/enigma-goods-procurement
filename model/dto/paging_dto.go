package dto

//parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

//return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

//response
type Paging struct {
	Page        int `json:"paging"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
