package util

import (
	"game/internal/logger"
	"os"
	"strconv"
)

// Getenv returns the environment
func Getenv(key string, def interface{}) interface{} {
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			logger.Debugf("%+v", err)
			os.Exit(1)
		}
	}()

	value := os.Getenv(key)

	if len(value) == 0 {
		return def
	} else {

		var val interface{}
		var err error

		switch def.(type) {
		case bool:
			val, err = strconv.ParseBool(value)
		case int:
			val, err = strconv.Atoi(value)
		case int64:
			val, err = strconv.ParseInt(value, 10, 64)
		case float32:
			val, err = strconv.ParseFloat(value, 32)
		case float64:
			val, err = strconv.ParseFloat(value, 64)
		default:
			val = value
		}

		if err != nil {
			panic(err)
		}

		return val
	}
}
