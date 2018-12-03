package tools

type ApiResponse struct{
	Code	int `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

const (
	SUCCESS = 10000
	ERROR = -10000
	UserNotFound = -10001
)

func Success(data interface{}, msg ...string) *ApiResponse {
	retMsg := "success"
	if (len(msg) > 0) {
		retMsg = msg[0]
	}
	ret := &ApiResponse{Code: SUCCESS, Data: data, Msg: retMsg}
	return ret
}

func Error(data interface{}, msg string, code int) *ApiResponse {
	ret := &ApiResponse{Code: code, Data: data, Msg: msg}
	return ret
}