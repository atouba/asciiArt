package main

// When there is --output flag, no terminal size checking
// should be done
// The characters are distorted when there are extra spaces

import (
	"flag"
	"fmt"
	"log"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/alignement"
	"01.gritlab.ax/git/atouba/ascii-art/basic"
	"01.gritlab.ax/git/atouba/ascii-art/color"
	"01.gritlab.ax/git/atouba/ascii-art/reverse"
)

func main() {
	align := flag.String("align", "", "specify alignment (left, center, right or justify)")
	clr := flag.String("color", "", "specify color (e.g. red, green, cyan etc.)")
	outputFileName := flag.String("output", "", "specify output file name")
	reverseFileName := flag.String("reverse", "", "specify name of the file that contains ascii art chars to be reversed")
	flag.Parse()

	if *align != "left" && *align != "center" && *align != "right" && *align != "justify" && *align != "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		os.Exit(1)
	}

	args := flag.Args()
	argsLength := len(args)

	if (argsLength == 0 && *reverseFileName == "") || argsLength > 3 {
		log.Fatal(`Usage:
      ./ascii_art <string>
      or
      ./ascii_art <string> <banner>
      or
      ./ascii_art . [OPTION] [STRING]`)
	}

	output := ""

  if *reverseFileName != "" {
    reverse.Reverse(*reverseFileName)
    return
  }
	if argsLength == 1 && *clr == "" {
		output = basic.Basic(args[0], "", *clr, "standard", *align)
	} else if argsLength == 2 && *clr == "" {
		output = basic.Basic(args[0], "", *clr, args[1], *align)
	} else {
		specifiedColor := *clr
		if argsLength == 1 {
			args = append(args, args[0]) // the whole string as substring to color
		}
		_, ok := color.Colors[specifiedColor]
		if !ok {
			log.Fatal("error: color doesn't exist")
		}
    output = basic.Basic(args[1], args[0], specifiedColor, "standard", *align)
	}
	output = alignement.AlignLCR(output, *align)

  // Output flag can be empty, should error. This condition should check if it wasn't typed at all.
  if *outputFileName != "" {
    f, err := os.OpenFile(*outputFileName, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
    if err != nil { log.Fatal(err) }
    f.WriteString(output)
  } else {
    fmt.Print(output)
  }
}

