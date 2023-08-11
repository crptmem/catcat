package main

import (
  "fmt"

  "github.com/spf13/viper"

)

func readConfig(path string) {
  fmt.Println("Reading configuration...")

  viper.SetConfigFile(path)

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Printf("Error reading config file: %s", err)
    return
  }
}

func getGameEntries() ([]interface{}) {
  GameEntries := viper.Get("GameEntries").([]interface{})
  return GameEntries
}

func getGameEntry(name string) (map[string]interface{}){
  var entries = getGameEntries()

  for _, entry := range entries {
		// Type-assert the entry to a map[string]interface{}
		entryMap := entry.(map[string]interface{})
    if entryMap["name"] == name {
      return entryMap
    }
	}
  return nil
}
