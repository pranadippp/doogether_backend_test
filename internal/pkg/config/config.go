package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
}

func (c *AppConfig) init(prefix string) {
	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = fmt.Sprintf("%s.%s", prefix, env)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fmt.Sprintf("%s.%s", env, "yaml"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func NewAppConfig() *AppConfig {
	c := AppConfig{}
	c.init("")

	return &c
}

func NewAppConfigWithPrefix(prefix string) *AppConfig {
	c := AppConfig{}
	c.init(prefix)

	return &c
}

func (c *AppConfig) GetString(key string) string {
	return viper.GetString(key)
}
