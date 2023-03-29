package responses

func CreateRespon(code int, data interface{}, status, message string) *Resp {
	return &Resp{Status: status, Code: code, Message: message, Data: data}
}
