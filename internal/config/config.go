package config

import (
	"os"

	"github.com/bytedance/sonic"
)

var Conf *Config

type Config struct {
	Db struct {
		Username string `json:"username"`
		Password string `json:"password"`
		DbName   string `json:"dbName"`
		DbHost   string `json:"dbHost"`
		DbPort   string `json:"dbPort"`
	} `json:"db"`
	Api struct {
		Addr     string `json:"addr"`
		LogLevel string `json:"logLevel"`
	}
}

type ConfigParser interface {
	Parse() *Config
}

type JsonConfigParser struct {
	path string
}

func NewJsonConfigParser(path string) ConfigParser {
	return &JsonConfigParser{path: path}
}

func (parser *JsonConfigParser) Parse() *Config {
	raw, err := os.ReadFile(parser.path)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := sonic.Unmarshal(raw, &config); err != nil {
		panic(err)
	}
	Conf = &config
	return &config
}
