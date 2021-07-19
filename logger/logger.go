package logger

func Trace(args ...interface{}) {
	Logger.Trace(args)
}
func Debug(args ...interface{}) {
	Logger.Debug(args)
}
func Info(args ...interface{}) {
	Logger.Info(args)
}
func Warn(args ...interface{}) {
	Logger.Warn(args)
}

func Error(args ...interface{}) {
	Logger.Error(args)
}
