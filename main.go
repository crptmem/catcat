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
  var setWine = flag.Bool("setwine", false, "Set Wine location")  
  
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
      os.Exit(0)
    }
  } else if *setWine {
    if flag.NArg() < 1 { 
      flag.Usage()
      os.Exit(1)
    } else {
      viper.Set("winelocation", flag.Arg(0))
      viper.WriteConfig()
      fmt.Println("Success")
      os.Exit(0)
    }
  } else {
    runInterfaceList()
  }
}

