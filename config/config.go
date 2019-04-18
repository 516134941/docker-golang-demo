package config

import (
	"github.com/BurntSushi/toml"
)

// Config 对应配置文件结构
type Config struct {
	Listen string `toml:"listen"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}

// GetListenAddr 监听地址
func (c Config) GetListenAddr() string {
	return c.Listen
}
