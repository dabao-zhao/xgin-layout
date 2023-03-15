package config

import "github.com/dabao-zhao/xgin/config"

type Config struct {
	Db   config.DbConfig   `toml:"db"`
	Http config.HttpConfig `toml:"http"`
	Log  config.LogConfig  `toml:"log"`
}
