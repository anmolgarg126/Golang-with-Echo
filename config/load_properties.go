package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Configuration Configurations

func setProperties() {
	fmt.Println("APP_ENV:", os.Getenv("APP_ENV"))

	os.Setenv("APP_ENV", "qa")

	os_env := os.Getenv("APP_ENV")
	fmt.Println("APP_ENV:", os_env)

	viper.SetConfigName(fmt.Sprintf("%s_config", os_env))
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	// viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	// fmt.Println("Reading variables using the model..")
	// fmt.Println("Database is\t", Configuration.Database.DBName)
	// fmt.Println("Port is\t\t", Configuration.Server.Port)
	// fmt.Println("EXAMPLE_PATH is\t", Configuration.EXAMPLE_PATH)
	// fmt.Println("EXAMPLE_VAR is\t", Configuration.EXAMPLE_VAR)

	// // Reading variables without using the model
	// fmt.Println("\nReading variables without using the model..")
	// fmt.Println("Database is\t", viper.GetString("database.dbname"))
	// fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	// fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	// fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
}
