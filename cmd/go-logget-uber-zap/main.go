package main

import (
	"fmt"

	c "github.com/BNPrashanth/go-logger-uber-zap/configs"

	"github.com/spf13/viper"
)

func main() {
	c.InitializeViper()
	fmt.Println(viper.GetString("logger-output-path"))
}
