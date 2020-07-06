package controller

type Response struct {
	Status string      `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

type List struct {
	Data  interface{} `json:"data"`
	Total interface{} `json:"total"`
	Size  interface{} `json:"size"`
	Page  interface{} `json:"page"`
}

func ReturnData(status string, data interface{}, msg string) (r *Response) {
	r = &Response{Status: status, Msg: msg, Data: data}
	return
}
