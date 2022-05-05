package config

type BaseResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewBaseResp(code int, msg string, data interface{}) *BaseResp {
	return &BaseResp{Code: code, Msg: msg, Data: data}
}
