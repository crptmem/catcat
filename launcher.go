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
    fmt.Printf("err: %s, path: %s", err, cmd.Path)
  }
}

func launchGame(path string, wrapper string) error {
  wineApp := viper.GetString("winelocation")

  if len(wrapper) == 0 {
    return nil
  }

  if _, err := os.Stat(path); os.IsNotExist(err) {
    return err
  }

  discordBridge := getConfigDir() + "/catcat/bin/winediscordipcbridge.exe"
  if _, err := os.Stat(discordBridge); os.IsNotExist(err) {
    // not critical, continue
  } else {
    bridgeCmd := exec.Command(wineApp, discordBridge)
    go execute(bridgeCmd)
  }

  cmd := exec.Command(wrapper, wineApp, path)
  go execute(cmd)
  return nil
}

