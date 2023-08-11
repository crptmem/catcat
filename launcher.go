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

  
  if len(wrapper) == 0 {
    cmd := exec.Command(wineApp, path)
    go execute(cmd)
  } else {
    cmd := exec.Command(wrapper, wineApp, path)
    go execute(cmd)
  }
  return nil
}

