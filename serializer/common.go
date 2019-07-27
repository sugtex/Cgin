package serializer

//基础回复序列器
type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data" `
	Msg string `json:"msg" `
	Error string `json:"error" `
}

