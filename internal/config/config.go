package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type IConfig interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetDuration(key string) time.Duration
}

var Cfg *viper.Viper

func newConfig() *viper.Viper {

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("json")

	cfg.AddConfigPath(getConfigPath())

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg.WatchConfig()

	return cfg
}

func getConfigPath() (path string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	slice := strings.Split(wd, "notification-service")
	path = slice[0] + "notification-service/internal/config"
	return
}

func ReadConfig() {
	Cfg = newConfig()
}
