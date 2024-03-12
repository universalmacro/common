package config

import (
	"github.com/spf13/viper"
	"github.com/universalmacro/common/singleton"
)

type Configuration interface {
	GetValue(key string) any
	SetValue(key string, value any) *Configuration
	GetConfig() map[string]any
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetFloat64(key string) float64
}

type MapConfig struct {
	config map[string]any
}

func (m *MapConfig) GetConfig() map[string]any {
	return m.config
}

func (m *MapConfig) SetValue(key string, value any) *MapConfig {
	m.config[key] = value
	return m
}

func (m *MapConfig) GetValue(key string) any {
	return m.config[key]
}

func (m *MapConfig) GetString(key string) string {
	return m.config[key].(string)
}

func (m *MapConfig) GetInt(key string) int {
	return m.config[key].(int)
}

func (m *MapConfig) GetBool(key string) bool {
	return m.config[key].(bool)
}

func (m *MapConfig) GetFloat64(key string) float64 {
	return m.config[key].(float64)
}

func NewMapConfig() *MapConfig {
	return &MapConfig{config: make(map[string]any)}
}

func NewSingletonMapConfig() func() *MapConfig {
	return singleton.EagerSingleton(func() *MapConfig {
		return NewMapConfig()
	})
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}

func GetString(s string) string {
	return viper.GetString(s)
}

func GetInt(s string) int {
	return viper.GetInt(s)
}

func GetBool(s string) bool {
	return viper.GetBool(s)
}

func GetFloat64(s string) float64 {
	return viper.GetFloat64(s)
}
