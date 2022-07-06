package contract

import "github.com/spf13/viper"

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetViper() *viper.Viper
}
