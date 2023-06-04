package models

import (
	"github.com/spf13/viper"
)

type Config struct {
	SqlServer SqlServer `yaml:"sqlserver"`
	Api       Api       `yaml:"api"`
	Jwt       JWT       `yaml:"jwt"`
}

type SqlServer struct {
	Server   string `yaml:"server"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Api struct {
	Port string `yaml:"port"`
}

type JWT struct {
	SecretKey string `yaml:"secret"`
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
