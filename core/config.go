package core

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

// Config export
type Config struct {
	DB struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
	}

	Jwt struct {
		AccessKey  string
		RefreshKey string
		Alg        string
	}
}

func _newInstance() *Config {
	c := Config{}
	dir, _ := os.Getwd()
	_ = configor.Load(&c, fmt.Sprintf("%s\\%s", dir, "config.json"))
	return &c
}

// GetConfig export
func GetConfig() *Config {
	return _newInstance()
}
