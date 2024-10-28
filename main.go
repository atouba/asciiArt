package main

import (
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
)

func main() {
	fmt.Print(callBasic(os.Args[1:]))
}

func callBasic(args []string) (out string) {
	if len(args) != 1 && len(args) != 2 {
		log.Fatal("Usage: go run . <text> or go run . <text> <style>")
	} else if len(args) == 2 {
		if args[1] != "standard" && args[1] != "shadow" && args[1] != "thinkertoy" {
			log.Fatal("Invalid style name. Use \"standard\", \"shadow\" or \"thinkertoy\"")
		}
		out += basic.Basic(args[0], args[1])
	} else {
		out += basic.Basic(args[0], "standard")
	}
	return out
}
