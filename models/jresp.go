package models

// JResp HTTP 请求的返回数据
type JResp struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            interface{}       `json:"body"`
}

// Gen 生成响应
func Gen(status int, body interface{}, contentType string) *JResp {
	j := JResp{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": contentType},
		Body:       body,
	}

	// bs, _ := json.Marshal(j)
	return &j
}
