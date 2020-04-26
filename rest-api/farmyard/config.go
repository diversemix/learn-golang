package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfiguration() {
	viper.SetDefault("BindAddress", "")
	viper.SetDefault("Port", 8080)

	viper.SetConfigName("config")          // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/farmyard/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.farmyard") // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig()            // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Could not find config file, so writing a fresh one.")
			err = viper.SafeWriteConfigAs("./config.yaml")

		}
	}
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
