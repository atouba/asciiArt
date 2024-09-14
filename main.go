package main

import (
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
	"01.gritlab.ax/git/atouba/ascii-art/color"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 4 {
		fmt.Println(`Error: usage:
		./ascii_art <string>
		or
		./ascii_art <string> <banner>`)
		return
	}
	if len(os.Args) == 2 {
		basic.Basic(os.Args[1], "standard")
	} else if len(os.Args) == 3 {
		basic.Basic(os.Args[1], os.Args[2])
	} else if len(os.Args) == 4 {
    specifiedColor := os.Args[1][8:]
    _, ok := color.Colors[specifiedColor]
    if !ok {
      log.Fatal("error: color doesn't exist")
    }
    color.Color(os.Args[3], os.Args[2], specifiedColor, "standard")
  }
}
