package config

import (
	"github.com/spf13/viper"
	"sync"
)

//const (
//	DB_HOST     = "db-hub.docker"
//	DB_PORT     = "3306"
//	DB_USERNAME = "mCn3-ilem-menoLogy"
//	DB_PASSWORD = "mCn3-ilem-menoLogy"
//	DB_NAME     = "evermos"
//)

var (
	config     *conf
	configOnce sync.Once
)

type Config interface {
	GetAppPort() string
	GetDbHost() string
	GetDbPort() string
	GetDbUsername() string
	GetDbPassword() string
	GetDbName() string
}

type conf struct {
	AppPort    string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
}

func (c *conf) GetAppPort() string {
	return c.AppPort
}
func (c *conf) GetDbHost() string {
	return c.DbHost
}

func (c *conf) GetDbPort() string {
	return c.DbPort
}

func (c *conf) GetDbUsername() string {
	return c.DbUsername
}

func (c *conf) GetDbPassword() string {
	return c.DbPassword
}

func (c *conf) GetDbName() string {
	return c.DbName
}

func NewConfig() Config {
	configOnce.Do(func() {
		v := viper.New()
		v.SetConfigFile(".env")
		v.AddConfigPath(".")

		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}

		config = &conf{}
		config.AppPort = v.GetString("APP_PORT")
		config.DbHost = v.GetString("MYSQL_HOSTNAME")
		config.DbPort = v.GetString("DB_MYSQL_PORT")
		config.DbUsername = v.GetString("MYSQL_USER")
		config.DbPassword = v.GetString("MYSQL_PASSWORD")
		config.DbName = v.GetString("MYSQL_DATABASE")

	})
	return config
}
