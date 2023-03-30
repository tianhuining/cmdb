package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

var (
	config *Config
)

func Conf()*Config{
	return config
}

func LoadTomlToConfig(filepath string) error{
	cfg := NewDefaultConfig()
	_,err :=toml.DecodeFile(filepath ,cfg)
	if err!=nil{
		return err
	}
	config = cfg
	return nil
}

func LoadEnvToConfig()error{
	cfg := NewDefaultConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	config = cfg
	return nil
}
