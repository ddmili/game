package config

import (
	"game/internal/util"

	_ "github.com/joho/godotenv/autoload"
)

// Config is the configuration
type Config struct {
	Debug    bool   `json:"debug"`
	Port     string `json:"port"`
	MysqlDSN string `json:"mysql_dsn"`
	Redis    dbConfig
}

type dbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

// Cfg holds the configuration
var Cfg *Config

// InitConfig initializes the config
func InitConfig() *Config {
	Cfg = &Config{}
	Cfg.Port = util.Getenv("GAME_PORT", ":8021").(string)
	Cfg.Debug = util.Getenv("GAME_DEBUG", true).(bool)
	Cfg.MysqlDSN = util.Getenv("MYSQL_DSN", "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local").(string)
	Cfg.Redis.Host = util.Getenv("REDIS_HOST", "127.0.0.1").(string)
	Cfg.Redis.Port = util.Getenv("REDIS_HOST", 6379).(int)
	Cfg.Redis.Password = util.Getenv("REDIS_PASSWORD", "").(string)
	return Cfg
}
