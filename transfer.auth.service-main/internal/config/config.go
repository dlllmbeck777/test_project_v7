package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/transferMVP/transfer.webapp/assets"
	"os"
)

var Config ConfigFaceCollector

type ConfigFaceCollector struct {
	Server         *ServerConfig   `json:"server"`
	GoogleSettings *GoogleSettings `json:"google_settings"`
	Pg             *Pg             `json:"pg"`
	Redis          *Redis          `json:"redis"`
}

type Pg struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DbName string `json:"db_name"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}
type Redis struct {
	Dsn string `json:"dsn"`
	Key string `json:"key"`
}

type ServerConfig struct {
	Addr    string `json:"addr"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type GoogleSettings struct {
	HostVerifyToken string `json:"host_verify_token"`
}

func Init() error {
	var flagEnv string
	flag.StringVar(&flagEnv, "m", "", "")
	flag.Parse()
	if flagEnv == "" {
		flagEnv = assets.DefaultFlag
	}

	file, err := os.ReadFile(fmt.Sprintf("config/config_%s.json", flagEnv))
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &Config)
}
