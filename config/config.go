package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	LogFile      string `json:"log_file"`
	AOFFile      string `json:"aof_file"`
	SnapshotFile string `json:"snapshot_file"`
	MaxConn      int    `json:"max_connections"`
	LogLevel     string `json:"log_level"`
}

func LoadConfig(file string) *Config {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
