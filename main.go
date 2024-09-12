package main

import (
	"fmt"
	"os"
	"github.com/atouba/ascii_art/basic"
	// "github.com/atouba/ascii_art/reverse"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println(`Error: usage:
		./ascii_art <string>
		or
		./ascii_art <string> <banner>`)
		return
	}
	if len(os.Args) == 2 || os.Args[2] == "standard" {
		basic.Basic(os.Args[1], "standard")
	} else {
		basic.Basic(os.Args[1], os.Args[2])
	}
	// reverse.Reverse(os.Args[1])

}
