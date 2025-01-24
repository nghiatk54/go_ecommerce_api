package initialize

import (
	"fmt"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	// Get config struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
}
