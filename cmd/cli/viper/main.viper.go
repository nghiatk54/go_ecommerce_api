package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
	// Set config
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	// Read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %v", err))
	}
	fmt.Println("Server port: ", viper.GetInt("server.port"))
	fmt.Println("Key jwt: ", viper.GetString("security.jwt.key"))
	// Get config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
	fmt.Println("Config port: ", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, Password: %s, Host: %s, DbName: %s\n", db.User, db.Password, db.Host, db.DbName)
	}

}
