package logger

import (
	"go.uber.org/zap"
)

var  l *zap.Logger


// InitLogger initializes the logger
func InitLogger(debug bool) {
	l = NewLogger(debug)
}


 // NewLogger creates a new logger
func NewLogger(debug bool) *zap.Logger {
	var l *zap.Logger
	if debug {
		l, _ = zap.NewDevelopment()
	}else{
		l, _ = zap.NewProduction()
	}
	return  l
}

// Infof formats the given message
func Infof(format string, args ...interface{}) {
	l.Sugar().Infof(format , args )
}


// Fatalf formats the given message
func Fatalf(format string, args ...interface{}) {
	l.Sugar().Fatalf(format,args)
}

// Warnf formats the given message
func Warnf(format string, args ...interface{}) {
	l.Sugar().Warnf(format,args)
}


// Errorf formats the given message
func Errorf(format string, args ...interface{}) {
	l.Sugar().Errorf(format,args)
}


// Debugf formats the given message
func Debugf(format string, args ...interface{}) {
	l.Sugar().Debugf(format,args)
}



