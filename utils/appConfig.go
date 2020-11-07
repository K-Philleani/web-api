package utils

import (
	"bufio"
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _acf *AppConfig


func ParseConfig(path string) (*AppConfig, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_acf)
	if err != nil {
		return nil, err
	}
	return _acf, nil
}