package config

import (
	"encoding/json"
	"fmt"
)

type ServerConfig struct {
	Name   string
	IpAddr string `json:"ip_addr"`
	Port   string `json:",omitempty"`
}

type GlobalConfig struct {
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

func Init(fileData []byte) error {
	conf := GlobalConfig{}

	if err := json.Unmarshal(fileData, &conf); err != nil {
		return err
	}

	for _, value := range conf.Server {
		fmt.Printf("%v\n", value)
	}
	return nil
}
