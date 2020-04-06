package config

import (
	"encoding/json"
)

type AppConfig struct {
	Logfile string
}

type ServerConfig struct {
	Name   string
	IpAddr string `json:"ip_addr"`
	User   string
	Port   string `json:",omitempty"`
}

type GlobalConfig struct {
	App    AppConfig
	Server []ServerConfig
}

func (s *ServerConfig) UnmarshalJSON(data []byte) error {
	type aliasServerConfig ServerConfig
	tmp := &aliasServerConfig{
		Port: "22",
	}

	_ = json.Unmarshal(data, tmp)

	*s = ServerConfig(*tmp)
	return nil
}

func Parse(fileData []byte) (*GlobalConfig, error) {
	conf := GlobalConfig{}
	if err := json.Unmarshal(fileData, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
