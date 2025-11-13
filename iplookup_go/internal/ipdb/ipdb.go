package ipdb

import (
	"errors"
	"net"
	"os"

	"github.com/ip2location/ip2location-go/v9"
	"iplookup/iplookup_go/internal/config"
	"iplookup/iplookup_go/internal/model"
)

// IPDB 封装IP数据库查询功能
type IPDB struct {
	ipv4DB *ip2location.DB
	ipv6DB *ip2location.DB
}

// Init 初始化IP数据库
func Init(cfg *config.Config) (*IPDB, error) {
	// 打开IPv4数据库
	ipv4, err := ip2location.OpenDB(cfg.IPDatabase.IPv4DB)
	if err != nil {
		return nil, errors.New("无法打开IPv4数据库: " + err.Error())
	}

	// 打开IPv6数据库
	ipv6, err := ip2location.OpenDB(cfg.IPDatabase.IPv6DB)
	if err != nil {
		ipv4.Close() // 关闭已打开的IPv4数据库
		return nil, errors.New("无法打开IPv6数据库: " + err.Error())
	}

	return &IPDB{
		ipv4DB: ipv4,
		ipv6DB: ipv6,
	}, nil
}

// Close 关闭数据库连接
func (ipdb *IPDB) Close() error {
	var err error
	
	if ipdb.ipv4DB != nil {
		if closeErr := ipdb.ipv4DB.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}
	
	if ipdb.ipv6DB != nil {
		if closeErr := ipdb.ipv6DB.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}
	
	return err
}

// QueryIPv4 查询IPv4地址信息
func (ipdb *IPDB) QueryIPv4(ipStr string) (model.IPv4Response, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil || ip.To4() == nil {
		return model.IPv4Response{
			Code:    1,
			Message: "无效的IPv4地址",
		}, errors.New("invalid ipv4 address")
	}

	result, err := ipdb.ipv4DB.Get_all(ipStr)
	if err != nil {
		return model.IPv4Response{
			Code:    2,
			Message: "查询失败: " + err.Error(),
		}, err
	}

	return model.IPv4Response{
		Code:    0,
		Message: "查询成功",
		Data: model.IPv4Info{
			IP:           ipStr,
			CountryCode:  result.Country_short,
			CountryName:  result.Country_long,
			Region:       result.Region,
			City:         result.City,
			Latitude:     result.Latitude,
			Longitude:    result.Longitude,
			ZipCode:      result.Zipcode,
			TimeZone:     result.Timezone,
			ISP:          result.Isp,
			Domain:       result.Domain,
			UsageType:    result.Usagetype,
			ASN:          result.Asn,
			ASName:       result.As,
		},
	}, nil
}

// QueryIPv6 查询IPv6地址信息
func (ipdb *IPDB) QueryIPv6(ipStr string) (model.IPv6Response, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil || ip.To16() == nil || ip.To4() != nil {
		return model.IPv6Response{
			Code:    1,
			Message: "无效的IPv6地址",
		}, errors.New("invalid ipv6 address")
	}

	result, err := ipdb.ipv6DB.Get_all(ipStr)
	if err != nil {
		return model.IPv6Response{
			Code:    2,
			Message: "查询失败: " + err.Error(),
		}, err
	}

	return model.IPv6Response{
		Code:    0,
		Message: "查询成功",
		Data: model.IPv6Info{
			IP:           ipStr,
			CountryCode:  result.Country_short,
			CountryName:  result.Country_long,
			Region:       result.Region,
			City:         result.City,
			Latitude:     result.Latitude,
			Longitude:    result.Longitude,
			ZipCode:      result.Zipcode,
			TimeZone:     result.Timezone,
			ISP:          result.Isp,
			ASN:          result.Asn,
			ASName:       result.As,
			Network:      result.Network,
		},
	}, nil
}

// GetIPType 判断IP类型
func (ipdb *IPDB) GetIPType(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "invalid"
	}
	if ip.To4() != nil {
		return "ipv4"
	}
	if ip.To16() != nil {
		return "ipv6"
	}
	return "unknown"
}