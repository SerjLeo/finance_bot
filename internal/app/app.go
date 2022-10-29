package app

import (
	"fmt"
	app "github.com/SerjLeo/finance_bot/internal/app/models"
)
import "github.com/spf13/viper"

func Run() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(err)
		return err
	}

	config := &app.Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return err
	}

	fmt.Printf("%+v", config)
	return nil
}
