package app

import (
	"github.com/SerjLeo/finance_bot/internal/service"
	"github.com/pkg/errors"
)
import "github.com/spf13/viper"

func Run() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "errors reading the config file")
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return errors.Wrap(err, "unable to decode into config struct")
	}

	serviceDeps := &service.Dependencies{}
	service := service.InitService(serviceDeps)

	err = InitRouter(service, config.Bot.Token)
	if err != nil {
		return errors.Wrap(err, "cant run bot")
	}

	return nil
}
