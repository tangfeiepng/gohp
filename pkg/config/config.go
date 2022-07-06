package config

import "github.com/spf13/viper"

type Config struct {
	viper *viper.Viper
}

func (config *Config) GetViper() *viper.Viper {
	return config.viper
}
func (config *Config) GetString(key string) string {
	return config.viper.GetString(key)
}
func (config *Config) GetBool(key string) bool {
	return config.viper.GetBool(key)
}
func (config *Config) GetInt(key string) int {
	return config.viper.GetInt(key)
}
