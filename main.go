package main

import (
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
)

func main() {
  if len(os.Args) != 2 {
    log.Fatal("Usage: go run . <string>")
  } else {
    fmt.Print(basic.Basic(os.Args[1], "standard"))
  }
}
