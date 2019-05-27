package main

import (
	c "github.com/BNPrashanth/go-logger-uber-zap/configs"
)

func main() {
	c.InitializeViper()
	c.InitializeZapCustomLogger()

	/*
		This is an example of a INFO level logger (generally considered as default log level).
			Has a higher priority than DEBUG level loggers, but lower than WARN level loggers.
			To get info level logger output, change the Level property in the Config struct in the initialization function as follows:
				> Level: zap.NewAtomicLevelAt(zapcore.InfoLevel)
	*/
	c.GenLog.Info("This is an INFO level message..")

	/*
		This is an example of a DEBUG level logger (Least priority loggers).
			To get debug level logger output, change the Level property in the Config struct in the initialization function as follows:
				> Level: zap.NewAtomicLevelAt(zapcore.DebugLevel)
	*/
	c.GenLog.Debug("This is a DEBUG level message..")

	// This is an example of a warning level logger
	c.GenLog.Warn("This is a WARNING level message..")

	// This is an example of a error level logger
	c.ErrLog.Error("This is an ERROR level message..")

	// This is an example of a dpanic level logger
	c.ErrLog.DPanic("This is a DPanic level message..")

	/*
		Panic and Fatal level loggers execute under any level and are considered to have the highest priority over all loggers,
		and if either of these levels are set for the logger, no others (level) logger messages will be given out.
		And if either of Panic or fatal level loggers execute, the application will exit.
			> With status of 1 for FATAL level
			> With status of 2 for PANIC level
	*/
	c.ErrLog.Panic("This is a PANIC level message..")
	c.ErrLog.Fatal("This is a FATAL level message..")
}
