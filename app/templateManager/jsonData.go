package templateManager

type JSONData struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}
