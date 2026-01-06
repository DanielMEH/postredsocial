package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

var (
	data string = "internal/infrastructure/data"
	path string = "appsettings.json"
)

type AppSettings struct {
	Log struct {
		FileName     string `json:"FileName"`
		ConsoleLevel string `json:"ConsoleLevel"`
	} `json:"log"`
	Config struct {
		Port       string `json:"Port"`
		Protocolo  string `json:"Protocolo"`
		Host       string `json:"Host"`
		Name       string `json:"Name"`
		Enviroment string `json:"Enviroment"`
	} `json:"Config"`

	RequestBodyLimit   string   `json:"RequestBodyLimit"`
	RequestParamsLimit string   `json:"RequestParamsLimit"`
	AllowMethods       []string `json:"AllowMethods"`
	AllowHeaders       []string `json:"AllowHeaders"`

	Cors struct {
		AllowOrigins []struct {
			Origin           string `json:"Origin"`
			AllowCredentials bool   `json:"AllowCredentials"`
		}
	} `json:"Cors"`
	NameDb string `json:"nameDb"`
}

func AppSettingsUnmarshalnFn() *AppSettings {
	originPath, err := filepath.Abs(".")

	if err != nil {

		return nil
	}
	jsonFilePath := filepath.Join(originPath, data, path)

	data, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil
	}

	var config AppSettings
	if err := json.Unmarshal(data, &config); err != nil {
		return nil
	}

	return &config
}
