package color

import (
	"fmt"
	"log"
	"os"
	"github.com/atouba/piscine"
)

var Colors = map[string]string {
  "Reset": "\033[0m",
  "red":  "\033[31m",
  "green": "\033[32m",
  "yellow":"\033[33m",
  "blue": "\033[34m",
  "purple":"\033[35m",
  "cyan": "\033[36m",
  "gray": "\033[37m",
  "white": "\033[97m",
}

func printMixedColor(str, subStr, color string) {
	asciiArtChars, err := os.ReadFile("./banners/standard.txt")
	if err != nil {
    log.Fatal("error reading banner file")
	}
	lines := piscine.Split(string(asciiArtChars), "\n")
  for iLine := range 8 {
    for i, char := range str {
      if char == ' ' {
        fmt.Print(lines[(int(char) - 32) * 8 + iLine][0:4])
      } else {
        if charInSubStr(str, i, subStr) { fmt.Print(Colors[color]) }
        fmt.Print(lines[(int(char) - 32) * 8 + iLine])
        if charInSubStr(str, i, subStr) { fmt.Print(Colors["Reset"]) }
      }
    }
    fmt.Println()
  }
}

func Color(str, subStr, color, banner string) {
  i := 0
  for ; i < len(str); {
    newLineIndex := index(str[i: ], "\\n")
    if newLineIndex == 0 {
      fmt.Println()
      i += 2
    } else {
      printMixedColor(str[i: i+newLineIndex], subStr, color)
      i += newLineIndex
    }
	}
}

func index(str, subStr string) int {
  iSubStr := piscine.Index(str, subStr)
  if iSubStr == -1 {
    return len(str)
  }
  return iSubStr
}

func charInSubStr(str string, i int, subStr string) bool {
  subStrIndex := piscine.Index(str, subStr)
  strIndex := subStrIndex
  for ; subStrIndex != -1;  {
    if i >= strIndex && i < strIndex + len(subStr) {
      return true
    }
    subStrIndex = piscine.Index(str[strIndex+1: ], subStr)+1
    if subStrIndex == 0 { break }
    strIndex += subStrIndex
  }
  return false
}

