package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//ServiceCfg represent whole service config
type ServiceCfg struct {
	Mysql MysqlCfg
}

//MysqlCfg represent Pulasr service config
type MysqlCfg struct {
	Server   string
	Port     string
	Table    string
	Database string
	User     string
	Passwd   string
}

//Config will get the value from path
func Config(path string) *ServiceCfg {
	// Config
	viper.SetConfigName("service") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv() // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	mysql := &MysqlCfg{}

	mysql.Server = viper.GetString("mysql.Server")
	mysql.Port = viper.GetString("mysql.Port")
	mysql.Table = viper.GetString("mysql.Table")
	mysql.Database = viper.GetString("mysql.Database")
	mysql.User = viper.GetString("mysql.User")
	mysql.Passwd = viper.GetString("mysql.Passwd")

	return &ServiceCfg{Mysql: *mysql}
}
