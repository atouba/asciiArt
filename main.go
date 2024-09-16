package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/alignement"
	"01.gritlab.ax/git/atouba/ascii-art/basic"
	"01.gritlab.ax/git/atouba/ascii-art/color"
	//"github.com/atouba/piscine"
)

func main() {
	align := flag.String("align", "", "specify alignment (left, center, right or justify)")
	clr := flag.String("color", "", "specify color (e.g. red, green, cyan etc.)")
	flag.Parse()

	if *align != "left" && *align != "center" && *align != "right" && *align != "justify" && *align != "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		os.Exit(1)
	}

	args := flag.Args()
	//args := os.Args[1:]

	argsLength := len(args)
	//fmt.Println(args, argsLength)

	if argsLength == 0 || argsLength > 3 {
		log.Fatal(`Usage:
      ./ascii_art <string>
      or
      ./ascii_art <string> <banner>
      or
      ./ascii_art . [OPTION] [STRING]`)
	}

	output := ""

	if argsLength == 1 && *clr == "" {
		//basic.Basic(args[0], "standard")
		output = basic.Basic(args[0], "standard", *align)

		//} else if argsLength == 2 && piscine.Index(args[0], "--color=") == -1 {
	} else if argsLength == 2 && *clr == "" {
		//basic.Basic(args[0], args[1])
		output = basic.Basic(args[0], args[1], *align)
	} else {

		//specifiedColor := args[0][8:]
		specifiedColor := *clr

		//if argsLength == 2 {
		if argsLength == 1 {
			args = append(args, args[0]) // the whole string as substring to color
		}
		_, ok := color.Colors[specifiedColor]
		if !ok {
			log.Fatal("error: color doesn't exist")
		}

		//color.Color(args[2], args[1], specifiedColor, "standard")
		output = color.Color(args[1], args[0], specifiedColor, "standard")
	}

	// align the text left, center or right
	output = alignement.AlignLCR(output, *align)

	fmt.Println(output)
}
