package basic

import (
	"fmt"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/alignement"
	"github.com/atouba/piscine"
)

func printBasic(str, banner, alignFlag string) string {
	asciiArtChars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return ""
	}

	out := ""

	lines := piscine.Split(string(asciiArtChars), "\n")
  leftSpacesLn, rightSpacesLn, insideSpacesLn := alignement.SpacesCount(str, alignFlag, banner)
	for iLine := range 8 {
    out += fmt.Sprint(alignement.SpacesString(leftSpacesLn))
		for _, char := range str {
			if char == ' ' {
        if insideSpacesLn > 0 {
          out += fmt.Sprint(alignement.SpacesString(insideSpacesLn))
        } else {
          out += fmt.Sprint(lines[(int(char)-32)*8+iLine])
        }
			} else {
				out += fmt.Sprint(lines[(int(char)-32)*8+iLine])
			}
		}
    out += fmt.Sprint(alignement.SpacesString(rightSpacesLn))
		out += fmt.Sprint("\n")
	}

	return out
}

func Basic(str, banner, alignFlag string) string {

	out := ""

	i := 0
	for ; i+1 < len(str) && str[i:i+2] == "\\n"; i += 2 {
		//fmt.Println()
		out += fmt.Sprint("\n")
	}
	for i < len(str) {
		stringLength := piscine.Index(str[i:], "\\n")
		if stringLength == -1 {
			stringLength = len(str[i:])
		}
		//printBasic(str[i:i+stringLength], banner)
		out += printBasic(str[i:i+stringLength], banner, alignFlag)

		// adding the length of the printed string
		i += stringLength
		for ; i+1 < len(str) && str[i:i+2] == "\\n"; i += 2 {
			//fmt.Println()
			out += fmt.Sprint("\n")
		}
	}

	return out
}
