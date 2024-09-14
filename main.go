package main

import (
	// "fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
	"01.gritlab.ax/git/atouba/ascii-art/color"
	"github.com/atouba/piscine"
)

func main() {
  args := os.Args[1: ]
  argsLength := len(args)

	if argsLength == 0 || argsLength > 3 {
    log.Fatal(`Usage:
      ./ascii_art <string>
      or
      ./ascii_art <string> <banner>
      or
      ./ascii_art . [OPTION] [STRING]`)
	}
	if argsLength == 1 {
		basic.Basic(args[0], "standard")
	} else if argsLength == 2 && piscine.Index(args[0], "--color=") == -1  {
		basic.Basic(args[0], args[1])
	} else {
    specifiedColor := args[0][8:]
    if argsLength == 2 { args = append(args, args[1]) }
    _, ok := color.Colors[specifiedColor]
    if !ok {
      log.Fatal("error: color doesn't exist")
    }
    color.Color(args[2], args[1], specifiedColor, "standard")
  }
}
