/*
Package configs implements configuration functio needed to run the application.
It contains the initialization functions to enable,
	Viper
	Uber-Zap-Custom-Logger
*/
package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
InitializeViper Funtion initializes viper to read configuration and/or environment variables in the application.
*/
func InitializeViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

/*
InitializeZapCustomLogger Funtion initializes a logger using uber-go/zap package in the application.
*/
func InitializeZapCustomLogger() {

}
