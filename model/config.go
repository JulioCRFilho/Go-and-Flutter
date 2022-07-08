package model

import "fmt"

type Config struct {
	Server    string `json:"server"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	Cluster   string `json:"cluster"`
	Host      string `json:"host"`
	Optionals string `json:"optionals"`
}

func (c Config) String() string {
	return fmt.Sprintf("%s://%s:%s@%s.%s/?%s", c.Server, c.User, c.Pass, c.Cluster, c.Host, c.Optionals)
}
