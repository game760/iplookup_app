package ipdb

import (
	"iplookup/iplookup_go/internal/config"
	"iplookup/iplookup_go/internal/database"
	"iplookup/iplookup_go/internal/model"
)

// IPDB IP数据库查询器
type IPDB struct {
	db     *database.DB
	cfg    *config.Config
	DBPath string
}

// Init 初始化IP查询器（修正为正确的初始化函数）
func Init(db *database.DB, cfg *config.Config) (*IPDB, error) {
	return &IPDB{
		db:     db,
		cfg:    cfg,
		DBPath: cfg.IPDatabase.IPv4Table, // 可以根据实际需求调整
	}, nil
}

// Close 关闭资源
func (db *IPDB) Close() error {
	// 实际项目中如果有需要关闭的资源可以在这里实现
	return nil
}

// Query 查询IP地理位置信息
func (db *IPDB) Query(ip string) model.IPQueryResponse {
	// 模拟无效IP校验
	if ip == "" {
		return model.IPQueryResponse{
			Code:    1,
			Message: "无效的IP地址",
			Data:    nil,
		}
	}

	// 模拟查询逻辑（实际项目中可以在这里实现数据库查询）
	location := &model.IPLocation{
		IP:        ip,
		Type:      "IPv4",
		Country:   "中国",
		Region:    "北京",
		City:      "北京",
		Latitude:  39.9042,
		Longitude: 116.4074,
	}

	// 返回成功响应
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