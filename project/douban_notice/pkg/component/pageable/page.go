package pageable

type PageVo struct {
	PageNo    int `json:"pageNo"`
	PageLimit int `json:"pageLimit"`
}

type PageRespVo struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
