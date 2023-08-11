package main
import (
   "os/exec"
   "fmt"

   "github.com/spf13/viper"
)

func launchGame(path string, wrapper string, discordPresense bool) {
  wineApp := viper.GetString("winelocation")
  cmd := exec.Command(wrapper, wineApp, path)
  err := cmd.Run()
  if err != nil {
    fmt.Println(err)
  }
}

