package model

type Pagination struct {
	Limit  int `json:"limit" form:"limit" uri:"limit" binding:"required"`
	Offset int `json:"offset" form:"offset" uri:"offset" binding:"required"`
}
