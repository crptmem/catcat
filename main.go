package main
import (
  "flag"
  "os"
  "fmt"

  "github.com/spf13/viper"
)


func main() {
  readConfig()
  var addGame = flag.Bool("addgame", false, "Add a new game (example: catcat --addgame name path wrapper)")  
  
  flag.Parse()

  if *addGame {
    if flag.NArg() < 3 { 
      flag.Usage()
      os.Exit(1)
    } else {
      var GameEntries = getGameEntries()
      
      newEntry := map[string]interface{}{"name": flag.Arg(0), "executablepath": flag.Arg(1), "wrappercommand": flag.Arg(2)}
      GameEntries = append(GameEntries, newEntry)
  
      viper.Set("gameentries", GameEntries)
      viper.WriteConfig()
      fmt.Println("Success")
    }
  } else {
    runInterfaceList()
  }
}

