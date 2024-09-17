package basic

import (
	"fmt"
	"os"

	"01.gritlab.ax/git/atouba/ascii-art/alignement"
	"01.gritlab.ax/git/atouba/ascii-art/color"
	"github.com/atouba/piscine"
)

type inputContent struct {
	str      string
	i        int // str's iterator index
	newLineI int // index of the nearest new line
	subStr   string
}

func clearCarReturns(s string) (out string) {
	for _, r := range s {
		if r != 13 {
			out += string(r)
		}
	}
	return
}

func printBasic(text *inputContent, banner, alignFlag, colorFl string) string {
	asciiArtChars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return ""
	}

	out := ""

	toLines := clearCarReturns(string(asciiArtChars))
	lines := piscine.Split(toLines, "\n")
	leftSpacesLn, rightSpacesLn, insideSpacesLn := alignement.SpacesCount(text.str[text.i:text.i+text.newLineI], alignFlag, banner)
	for iLine := range 8 {
		out += fmt.Sprint(alignement.SpacesString(leftSpacesLn))
		for i, char := range text.str[text.i : text.i+text.newLineI] {
			toAdd := lines[(int(char)-32)*8+iLine]
			if char == ' ' {
				if insideSpacesLn > 0 {
					out += fmt.Sprint(alignement.SpacesString(insideSpacesLn))
				} else {
					out += fmt.Sprint(toAdd)
				}
			} else {
				if color.CharInSubStr(text.str, text.i+i, text.subStr) {
					out += fmt.Sprint(color.Colors[colorFl])
				}
				out += fmt.Sprint(toAdd)
				if color.CharInSubStr(text.str, text.i+i, text.subStr) {
					out += fmt.Sprint(color.Colors["Reset"])
				}
			}
		}
		out += fmt.Sprint(alignement.SpacesString(rightSpacesLn))
		out += fmt.Sprint("\n")
	}

	return out
}

func Basic(str, subStr, colorFl, banner, alignFlag string) string {
	text := inputContent{}
	text.str = str
	text.subStr = subStr
	text.i = 0
	out := ""
	for text.i < len(text.str) {
		text.newLineI = color.Index(text.str[text.i:], "\\n")
		if text.newLineI == 0 {
			out += fmt.Sprintln()
			text.i += 2
		} else {
			out += printBasic(&text, banner, alignFlag, colorFl)
			text.i += text.newLineI
		}
	}
	return out
}

// func Basic(str, banner, alignFlag string) string {

// 	out := ""

// 	i := 0
// 	for i < len(str) {
// 		stringLength := piscine.Index(str[i:], "\\n")
// 		if stringLength == -1 {
// 			stringLength = len(str[i:])
// 		}
// 		//printBasic(str[i:i+stringLength], banner)
// 		out += printBasic(str[i:i+stringLength], banner, alignFlag)

// 		// adding the length of the printed string
// 		i += stringLength
// 		for ; i+1 < len(str) && str[i:i+2] == "\\n"; i += 2 {
// 			//fmt.Println()
// 			out += fmt.Sprint("\n")
// 		}
// 	}

// 	return out
// }
