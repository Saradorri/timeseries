package app

import (
	"edgecom.ai/timeseries/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func (a *application) setupViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		msg := fmt.Sprintf("viper read config error: %s", err.Error())
		return errors.New(msg)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var c utils.ServiceConfig
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	a.config = &c
	return nil
}
