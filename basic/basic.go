package basic

import (
	"fmt"
	"os"

	"github.com/atouba/piscine"
)

func printBasic(str, banner string) string {
	asciiArtChars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return ""
	}

	out := ""

	lines := piscine.Split(string(asciiArtChars), "\n")
	for iLine := range 8 {
		for _, char := range str {
			if char == ' ' {
				//fmt.Print(lines[(int(char) - 32) * 8 + iLine][0:4])
				out += fmt.Sprint(lines[(int(char)-32)*8+iLine])
			} else {
				//fmt.Print(lines[(int(char) - 32) * 8 + iLine])
				out += fmt.Sprint(lines[(int(char)-32)*8+iLine])
			}
		}
		//fmt.Println()
		out += fmt.Sprint("\n")
	}

	return out
}

func Basic(str, banner string) string {

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
		out += printBasic(str[i:i+stringLength], banner)

		// adding the length of the printed string
		i += stringLength
		for ; i+1 < len(str) && str[i:i+2] == "\\n"; i += 2 {
			//fmt.Println()
			out += fmt.Sprint("\n")
		}
	}

	return out
}
