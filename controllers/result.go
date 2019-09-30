package controllers

type Result struct {
	Status bool		   `json:"status"`
	Mag    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiResult(status bool, msg interface{}, data interface{}) (result *Result) {
	result = &Result{Status:status, Mag:msg, Data:data}
	return
}
