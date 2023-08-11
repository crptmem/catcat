package main

import (
  "fmt"
  "os"
  "path/filepath"

  "github.com/spf13/viper"
)

func getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, "/.config")
}


// Check is config present in $HOME/.config/catcat, if not - create.

func isConfigPresent() string {
	// Set the config file name and path
	configFileName := "config.json"
	configFilePath := filepath.Join(getConfigDir(), "catcat")

	// Set the config type (optional, default is JSON)
	viper.SetConfigType("json")

	// Set the config file name and path
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)

  // Populate config
  viper.SetDefault("gameentries", []interface{}{})

  viper.SetDefault("winelocation", "/usr/bin/wine")

	// Check if the config file already exists
	if _, err := os.Stat(filepath.Join(configFilePath, configFileName)); os.IsNotExist(err) {
		// Create the config directory if it doesn't exist
		if err := os.MkdirAll(configFilePath, 0755); err != nil {
			fmt.Println("Failed to create config directory:", err)
			return ""
		}

		// Create a new config file with default values
		if err := viper.SafeWriteConfigAs(filepath.Join(configFilePath, configFileName)); err != nil {
			fmt.Println("Failed to create config file:", err)
			return ""
		}

		fmt.Println("Config file created at", filepath.Join(configFilePath, configFileName))
	} else {
		// Read the existing config file
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Failed to read config file:", err)
			return ""
    }

	}

  return filepath.Join(configFilePath, configFileName)
}

func readConfig() {
  fmt.Println("Reading configuration...")
  viper.SetConfigFile(isConfigPresent())

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Printf("Error reading config file: %s", err)
    return
  }
}

func getGameEntries() ([]interface{}) {
  GameEntries := viper.Get("gameentries").([]interface{})
  return GameEntries
}

func getGameEntry(name string) (map[string]interface{}){
  var entries = getGameEntries()

  for _, entry := range entries {
    entryMap := entry.(map[string]interface{})
    if entryMap["name"] == name {
      return entryMap
    }
  }
  return nil
}
