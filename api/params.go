package api

type ListParams struct {
	Per  int    `json:"per" form:"per"`
	Page int    `json:"page" form:"page"`
	Key  string `json:"key" form:"key"`
}

type IdParams struct {
	Id int64 `json:"id" form:"id"`
}

type IdStrParams struct {
	Id string `json:"id" form:"id"`
}
