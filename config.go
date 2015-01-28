package config

import (
	"encoding/json"
	"os"
)

// Maps service names (i.e. "phil" or "collins") to their marathon app ID (i.e. /app/sre/sys/phil)
type ServiceAppMap map[string]string

type Config struct {
	// localhost
	MarathonHost string `json:"marathon-host,omitempty"`
	// 8080
	MarathonPort int `json:"marathon-port"`
	// service.ewr01.tumblr.net
	ServiceDomain string `json:"service-domain"`
	// {"phil":"/internal/sre/sys/phil","collins":"/internal/sre/sys/collins"}
	Services ServiceAppMap `json:"services,omitempty"`
}

func LoadConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	c := Config{}
	err = json.NewDecoder(f).Decode(&c)
	if c.MarathonPort == 0 {
		c.MarathonPort = 8080
	}
	return &c, nil
}
