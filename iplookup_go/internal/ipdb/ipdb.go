package ipdb

import (
	"iplookup/iplookup_go/internal/model"
)

// IPDB IP数据库查询器
type IPDB struct {
	DBPath string
}

// NewIPDB 初始化IP数据库实例
func NewIPDB(dbPath string) *IPDB {
	return &IPDB{
		DBPath: dbPath,
	}
}

// Close 关闭IP数据库资源（空实现，可根据实际需求扩展）
func (db *IPDB) Close() error {
	// 若有需要释放的资源（如文件句柄），可在此实现
	return nil
}

// Query 查询IP地理位置信息
func (db *IPDB) Query(ip string) model.IPQueryResponse {
	if ip == "" {
		return model.IPQueryResponse{
			Code:    1,
			Message: "无效的IP地址",
			Data:    nil,
		}
	}

	location := &model.IPLocation{
		IP:        ip,
		Type:      "IPv4",
		Country:   "中国",
		Region:    "北京",
		City:      "北京",
		Latitude:  39.9042,
		Longitude: 116.4074,
	}

	return model.IPQueryResponse{
		Code:    0,
		Message: "查询成功",
		Data:    location,
	}
}

// BatchQuery 批量查询IP
func (db *IPDB) BatchQuery(ips []string) model.IPQueryResponse {
	if len(ips) == 0 {
		return model.IPQueryResponse{
			Code:    1,
			Message: "IP列表不能为空",
			Data:    nil,
		}
	}

	var locations []*model.IPLocation
	for _, ip := range ips {
		locations = append(locations, &model.IPLocation{
			IP:      ip,
			Country: "中国",
		})
	}

	return model.IPQueryResponse{
		Code:    0,
		Message: "批量查询成功",
		Data:    locations,
	}
}