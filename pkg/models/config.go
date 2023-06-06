package models

import (
	"github.com/spf13/viper"
)

type Config struct {
	SqlServer SqlServer `yaml:"sqlserver"`
	Api       Api       `yaml:"api"`
	Jwt       JWT       `yaml:"jwt"`
	Smtp      Smtp      `yaml:"smtp"`
}

type SqlServer struct {
	Server   string `yaml:"server"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Api struct {
	Port    string   `yaml:"port"`
	AppKeys []string `yaml:"appkeys"`
}

type JWT struct {
	SecretKey string `yaml:"secret"`
}

type Smtp struct {
	Host     string `yaml:"host"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

func (config *Config) GetConfigValues() {
	v := viper.New()
	v.SetTypeByDefaultValue(true)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./resources")
	err := v.ReadInConfig()
	if err != nil {
		panic("Config not found")
	}
	sub := v.Sub("local")
	unMarshallErr := sub.Unmarshal(config)

	if unMarshallErr != nil {
		panic("Unmarshall error")
	}
}
