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

// isJustSpacesBefore checks if there are only spaces in a string str before index i
func isJustSpacesBefore(str string, i int) bool {
	for ; i >= 0; i-- {
		if str[i] != ' ' {
			return false
		}
	}
	return true
}

// isJustSpacesAfter checks if there are only spaces in a string str after index i
func isJustSpacesAfter(str string, i int) bool {
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			return false
		}
	}
	return true
}

// printBasic generates ascii art for the function Basic()
func printBasic(text *inputContent, banner, alignFlag, colorFl string) string {
	asciiArtChars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return ""
	}

	out := ""

	lines := piscine.Split(alignement.ClearCarReturns(string(asciiArtChars)), "\n")
	inLineStr := text.str[text.i : text.i+text.newLineI]
	leftSpacesLn, rightSpacesLn, insideSpacesLn := alignement.SpacesCount(inLineStr, alignFlag, banner)
	for iLine := range 8 {
		out += fmt.Sprint(alignement.SpacesString(leftSpacesLn))
		for i, char := range inLineStr {
			toAdd := lines[(int(char)-32)*8+iLine]
			if char == ' ' {
				if insideSpacesLn > 0 {
					if isJustSpacesBefore(inLineStr, i) || isJustSpacesAfter(inLineStr, i) ||
						(i+1 < len(inLineStr) && inLineStr[i+1] == ' ') {
						continue
					}
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

// Basic returns an ascii art text string from a string str according to possible flags and a chosen style
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
