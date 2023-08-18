package config

import (
	"fmt"
	"os"

	"github.com/septian03yogi/utils/common"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type FileConfig struct {
	FilePath string
}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
}

func (cfg *Config) ReadConfig() (err error) {
	err = common.LoadEnv()
	if err != nil {
		return err
	}

	cfg.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	cfg.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	cfg.FileConfig = FileConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}

	if cfg.DbConfig.Host == "" || cfg.DbConfig.Port == "" || cfg.DbConfig.Name == "" || cfg.DbConfig.User == "" || cfg.DbConfig.Password == "" || cfg.ApiConfig.ApiHost == "" || cfg.ApiConfig.ApiPort == "" {
		return fmt.Errorf("missing required environment variable")
	}

	return nil
}

func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}
	err = cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return
}
