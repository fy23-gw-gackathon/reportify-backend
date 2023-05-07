package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Database     `yaml:"database"`
	Datastore    `yaml:"datastore"`
	App          `yaml:"app"`
	OpenAI       `yaml:"openAI"`
	AllowOrigins []string `yaml:"allowOrigins"`
	Cognito      `yaml:"cognito"`
}

type Datastore struct {
	Address string `yaml:"address"`
}

type App struct {
	Debug bool `yaml:"debug"`
	Port  int  `yaml:"port"`
}

type Cognito struct {
	ClientID   string `yaml:"clientID"`
	UserPoolID string `yaml:"userPoolID"`
}

type Database struct {
	Url string `yaml:"url"`
}

type OpenAI struct {
	Key string `yaml:"key"`
}

func Load() Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")

	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var conf Config

	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalln(err)
	}

	return conf
}
