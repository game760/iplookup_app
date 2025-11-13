package model

// IPv4Info IPv4详细信息
type IPv4Info struct {
	IP           string  `json:"ip"`
	CountryName  string  `json:"country_name"`
	Region       string  `json:"region"`       // 地区
	Province     string  `json:"province"`     // 新增省份
	City         string  `json:"city"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ISP          string  `json:"isp,omitempty"` // 运营商
}

// IPv4Response IPv4查询响应
type IPv4Response struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    IPv4Info  `json:"data,omitempty"`
}

// IPv6Info IPv6详细信息
type IPv6Info struct {
	IP           string  `json:"ip"`
	CountryName  string  `json:"country_name"`
	Region       string  `json:"region"`
	Province     string  `json:"province"`
	City         string  `json:"city"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ISP          string  `json:"isp,omitempty"`
}

// IPv6Response IPv6查询响应
type IPv6Response struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    IPv6Info  `json:"data,omitempty"`
}

// ErrorResponse 错误响应
func ErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    1,
		"message": message,
	}
}

// SuccessResponse 成功响应
func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    data,
	}
}