package main
import (
   "os/exec"
   "os"
   "fmt"

   "github.com/spf13/viper"
)

func execute(cmd *exec.Cmd){
  err := cmd.Run()
  if err != nil {
    fmt.Println(err)
  }
}

func launchGame(path string, wrapper string, discordPresence bool) error {
  wineApp := viper.GetString("winelocation")

  if _, err := os.Stat(path); os.IsNotExist(err) {
    return err
  }

  if discordPresence {
    discordBridge := getConfigDir() + "/catcat/bin/winediscordipcbridge.exe"
    if _, err := os.Stat(discordBridge); os.IsNotExist(err) {
      // not critical, continue
    } else {
      bridgeCmd := exec.Command(wineApp, discordBridge)
      go execute(bridgeCmd)
    }
  }

  cmd := exec.Command(wrapper, wineApp, path)
  go execute(cmd)
  return nil
}

