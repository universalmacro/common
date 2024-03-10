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

func (this *MapConfig) GetConfig() map[string]any {
	return this.config
}

func (this *MapConfig) SetValue(key string, value any) *MapConfig {
	this.config[key] = value
	return this
}

func (this *MapConfig) GetValue(key string) any {
	return this.config[key]
}

func (this *MapConfig) GetString(key string) string {
	return this.config[key].(string)
}

func (this *MapConfig) GetInt(key string) int {
	return this.config[key].(int)
}

func (this *MapConfig) GetBool(key string) bool {
	return this.config[key].(bool)
}

func (this *MapConfig) GetFloat64(key string) float64 {
	return this.config[key].(float64)
}

func NewMapConfig() *MapConfig {
	return &MapConfig{config: make(map[string]any)}
}

func NewSingletonMapConfig() singleton.Singleton[MapConfig] {
	return singleton.SingletonFactory(NewMapConfig, singleton.Lazy)
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
