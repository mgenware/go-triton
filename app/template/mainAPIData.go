package template

type MainAPIData struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
