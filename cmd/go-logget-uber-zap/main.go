package main

import (
	c "github.com/BNPrashanth/go-logger-uber-zap/configs"
)

func main() {
	c.InitializeViper()
	c.InitializeZapCustomLogger()

	c.Log.Info("This is an INFO level message..")
	c.Log.Debug("This is a DEBUG level message..")
	c.Log.Warn("This is a WARNING level message..")
	c.Log.Error("This is an ERROR level message..")
	c.Log.Panic("This is a PANIC level message..")
	c.Log.DPanic("This is a DPanic level message..")
	c.Log.Fatal("This is a FATAL level message..")
}
