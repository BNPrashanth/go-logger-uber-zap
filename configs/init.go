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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// GenLog variable is a globally accessible variable which will be initialized when the InitializeZapCustomLogger function is executed successfully.
	GenLog *zap.Logger

	// ErrLog variable is a globally accessible variable which will be initialized when the InitializeZapCustomLogger function is executed successfully.
	ErrLog *zap.Logger
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
	GenLog = initializeGeneralLogger()
	ErrLog = initializeErrorLogger()
	GenLog.Info("Logger Initialized..")
}

func initializeGeneralLogger() *zap.Logger {
	conf := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{viper.GetString("logger-output-path"), "stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:      "log",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			MessageKey:   "msg",
			EncodeName:   zapcore.FullNameEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	Log, _ := conf.Build()
	return Log
}

func initializeErrorLogger() *zap.Logger {
	conf := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.ErrorLevel),
		OutputPaths: []string{viper.GetString("logger-output-path"), "stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:      "log",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			MessageKey:   "msg",
			EncodeName:   zapcore.FullNameEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	Log, _ := conf.Build()
	return Log
}
