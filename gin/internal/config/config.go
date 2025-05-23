package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DB     DatabaseConfig `env-prefix:"PG_"`
	Server ServerConfig   `env-prefix:"HTTP_"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
}

type ServerConfig struct {
	Port string `env:"PORT"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
