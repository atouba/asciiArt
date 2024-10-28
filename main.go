package main

import (
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatal("Usage: go run . <text> or go run . <text> <style>")
	} else if len(os.Args) == 3 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			log.Fatal("Invalid style name. Use \"standard\", \"shadow\" or \"thinkertoy\"")
		}
		fmt.Print(basic.Basic(os.Args[1], os.Args[2]))
	} else {
		fmt.Print(basic.Basic(os.Args[1], "standard"))
	}
}
