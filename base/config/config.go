package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var Mega *Config

type LogCfg struct {
	Level      int    `json:"level"`
	Path       string `json:"path"`
	FileCount  int    `json:"fileCount"`
	StackTrace bool   `json:"stackTrace"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
	Console    bool   `json:"console"`
}

type HttpCfg struct {
	Port  uint16 `json:"port"`
	Debug bool   `json:"debug"`
}

type DbCfg struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	TablePrefix string `json:"tablePrefix"`
}

type SessionCfg struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
	TTL    int64  `json:"ttl"`
}

type Config struct {
	Http    *HttpCfg
	Log     *LogCfg
	Db      *DbCfg
	Session *SessionCfg
}

func init() {
	cfgPath := flag.String("conf", "mega.json", "-conf filepath")
	f, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		panic(fmt.Sprintf("load config file err: %v", err))
	}

	Mega = new(Config)
	if err = json.Unmarshal(f, Mega); err != nil {
		panic(fmt.Sprintf("load config file err: %v", err))
	}
}
