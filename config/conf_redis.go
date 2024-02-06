package config

import "fmt"

type Redis struct {
	Ip       string `yaml:"ip" json:"ip"`
	Password string `yaml:"password" json:"password"`
	Port     int    `yaml:"port" json:"port"`
	PoolSize int    `yaml:"pool_size" json:"poolSize"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Ip, r.Port)
}
