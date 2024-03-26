// SPDX-License-Identifier: Apache-2.0
package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Seed               string `json:"seed"`
	WsApiUrl           string `json:"ws_api_url"`
	HttpApiUrl         string `json:"http_api_url"`
	AppID              int    `json:"app_id"`
	Timeout            int    `json:"timeout"`
}

func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil
}
