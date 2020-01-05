package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var Mega *Config

type LogCfg struct {
	Level int `json:"level"`
	Path string `json:"path"`
	FileCount int `json:"fileCount"`
	StackTrace bool   `json:"stackTrace"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
	Console    bool   `json:"console"`
}

type HttpCfg struct {
	Port uint16 `json:"port"`
	Debug bool `json:"debug"`
}

type Config struct {
	Http *HttpCfg
	Log *LogCfg
}

func Init() error {
	cfgPath := flag.String("conf", "mega.json", "-conf filepath")
	f, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		return err
	}

	Mega = new(Config)
	if err = json.Unmarshal(f, Mega); err != nil {
		return err
	}

	return nil
}