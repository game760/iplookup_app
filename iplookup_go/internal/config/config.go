package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

// Config 全局配置结构体
type Config struct {
	Server struct {
		Port         string `yml:"port" yaml:"port"`
		ReadTimeout  int    `yml:"read_timeout" yaml:"read_timeout"`
		WriteTimeout int    `yml:"write_timeout" yaml:"write_timeout"`
	} `yml:"server" yaml:"server"`

	IPDatabase struct {
		IPv4DB string `yml:"ipv4_db" yaml:"ipv4_db"`
		IPv6DB string `yml:"ipv6_db" yaml:"ipv6_db"`
	} `yml:"ip_database" yaml:"ip_database"`

	API struct {
		Prefix          string   `yml:"prefix" yaml:"prefix"`
		JWTSecret       string   `yml:"jwt_secret" yaml:"jwt_secret"`
		RateLimit       int      `yml:"rate_limit" yaml:"rate_limit"`
		AllowOrigins    []string `yml:"allow_origins" yaml:"allow_origins"`
	} `yml:"api" yaml:"api"`

	JWT struct {
		Secret     string `yml:"secret" yaml:"secret"`
		ExpiresHours int  `yml:"expires_hours" yaml:"expires_hours"`
	} `yml:"jwt" yaml:"jwt"`
}

// Load 从YML/YAML文件加载配置，支持两种标签格式（优先yml）
func Load(path string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	// 解析YAML内容
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// 设置完整默认值
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Server.ReadTimeout == 0 {
		cfg.Server.ReadTimeout = 30
	}
	if cfg.Server.WriteTimeout == 0 {
		cfg.Server.WriteTimeout = 30
	}

	if cfg.IPDatabase.IPv4DB == "" {
		cfg.IPDatabase.IPv4DB = "./ipdata/ipv4.xdb"
	}
	if cfg.IPDatabase.IPv6DB == "" {
		cfg.IPDatabase.IPv6DB = "./ipdata/ipv6.xdb"
	}

	if cfg.API.Prefix == "" {
		cfg.API.Prefix = "/api/v1"
	}
	if cfg.API.RateLimit == 0 {
		cfg.API.RateLimit = 100
	}
	if len(cfg.API.AllowOrigins) == 0 {
		cfg.API.AllowOrigins = []string{"*"}
	}

	if cfg.JWT.Secret == "" {
		cfg.JWT.Secret = "default-secret-key-change-in-production"
	}
	if cfg.JWT.ExpiresHours == 0 {
		cfg.JWT.ExpiresHours = 24
	}

	return &cfg, nil
}