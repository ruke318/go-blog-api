package tools

type ApiResponse struct{
	Code   interface{} `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

const (
	SUCCESS = 10000
	ERROR = -10000
	UserNotFound = -10001
	UserAddFail = -10002
)

func Success(data interface{}, msg ...string) *ApiResponse {
	retMsg := "success"
	if len(msg) > 0 {
		retMsg = msg[0]
	}
	ret := &ApiResponse{Code: SUCCESS, Data: data, Msg: retMsg}
	return ret
}

func Error(msg ...interface{}) *ApiResponse {
	msgLen := len(msg)
	var code interface{}
	var	data interface{}
	var content interface{}
	if msgLen == 0 {
		code = ERROR
		content = "error request"
		data = ""
	} else if msgLen == 1 {
		code = ERROR
		content = msg[0]
		data = ""
	} else if msgLen == 2 {
		code = msg[1]
		content = msg[0]
		data = ""
	} else if msgLen == 3 {
		code = msg[1]
		content = msg[0]
		data = msg[2]
	} else {
		code = ERROR
		content = "error request"
		data = ""
	}
	ret := &ApiResponse{Code: code, Data: data, Msg: content}
	return ret
}