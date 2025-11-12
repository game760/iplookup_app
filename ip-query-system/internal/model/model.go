package model

// IPLocation IP地理位置信息
type IPLocation struct {
	IP        string  `json:"ip"`
	Type      string  `json:"type"`
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// IPInfo 详细IP信息（扩展字段）
type IPInfo struct {
	IP        string  `json:"ip"`
	Type      string  `json:"type"`
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ZipCode   string  `json:"zip_code"`
	Timezone  string  `json:"timezone"`
}

// IPQueryResponse IP查询响应结构
type IPQueryResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse 通用成功响应
func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "success",
		"Data":    data, // 注意字段名大小写需与前端一致
	}
}

// ErrorResponse 通用错误响应
func ErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    1,
		"message": message,
	}
}