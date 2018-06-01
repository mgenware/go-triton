package template

type JSONData struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
