package config

import "github.com/ilyakaznacheev/cleanenv"

type ServiceConfig struct {
	Port     string   `yaml:"port" env:"PORT" env-default:"3000"`
	Host     string   `yaml:"host" env:"HOST" env-default:"localhost"`
	Mode     string   `yaml:"mode" env:"MODE" env-default:"dev"`
	DBConfig DBConfig `yaml:"db_config"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

var cfg ServiceConfig

func GetConfig() (*ServiceConfig, error) {
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
