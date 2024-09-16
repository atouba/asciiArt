package color

import (
	// "fmt"
	// "log"
	// "os"

	"github.com/atouba/piscine"
)


var Colors = map[string]string{
	"Reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"gray":   "\033[37m",
	"white":  "\033[97m",
}

// func printMixedColor(text *inputContent, color string) string {
// 	asciiArtChars, err := os.ReadFile("./banners/standard.txt")
// 	if err != nil {
// 		log.Fatal("error reading banner file")
// 	}

// 	lines := piscine.Split(string(asciiArtChars), "\n")
// 	str := text.str[text.i : text.i+text.newLineI]
// 	out := ""
// 	for iLine := range 8 {
// 		for i, char := range str {
// 			if char == ' ' {

// 				//fmt.Print(lines[(int(char) - 32) * 8 + iLine][0:4])
// 				out += fmt.Sprint(lines[(int(char)-32)*8+iLine][0:4])
// 			} else {
// 				if CharInSubStr(text.str, text.i+i, text.subStr) {

// 					//fmt.Print(Colors[color])
// 					out += fmt.Sprint(Colors[color])
// 				}

// 				//fmt.Print(lines[(int(char) - 32) * 8 + iLine])
// 				out += fmt.Sprint(lines[(int(char)-32)*8+iLine])
// 				if CharInSubStr(text.str, text.i+i, text.subStr) {

// 					//fmt.Print(Colors["Reset"])
// 					out += fmt.Sprint(Colors["Reset"])
// 				}
// 			}
// 		}
// 		//fmt.Println()
// 		out += fmt.Sprintf("\n")
// 	}

// 	return out
// }

// func Color(str, subStr, color, banner string) string {
// 	text := InputContent{}
// 	text.str = str
// 	text.subStr = subStr
// 	text.i = 0
// 	out := ""
// 	for text.i < len(text.str) {
// 		text.newLineI = index(text.str[text.i:], "\\n")
// 		if text.newLineI == 0 {
// 			//fmt.Println()
// 			out += fmt.Sprint("\n")
// 			text.i += 2
// 		} else {
// 			//printMixedColor(&text, color)
// 			out += printMixedColor(&text, color)
// 			text.i += text.newLineI
// 		}
// 	}
// }
// 	return out

func Index(str, subStr string) int {
	iSubStr := piscine.Index(str, subStr)
	if iSubStr == -1 {
		return len(str)
	}
	return iSubStr
}

func CharInSubStr(str string, i int, subStr string) bool {
  if subStr == "" { return false }
	subStrIndex := piscine.Index(str, subStr)
	strIndex := subStrIndex
	for subStrIndex != -1 {
		if i >= strIndex && i < strIndex+len(subStr) {
			return true
		} else if i < strIndex {
			return false
		}
		subStrIndex = piscine.Index(str[strIndex+1:], subStr) + 1
		if subStrIndex == 0 {
			break
		}
		strIndex += subStrIndex
	}
	return false
}
