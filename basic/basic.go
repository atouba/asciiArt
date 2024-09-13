package basic

import (
	"fmt"
	"os"
	"github.com/atouba/piscine"
)

func printBasic(str, banner string) {
	asciiArtChars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}
	lines := piscine.Split(string(asciiArtChars), "\n")
  for iLine := range 8 {
    for _, char := range str {
      if char == ' ' {
        fmt.Print(lines[(int(char) - 32) * 8 + iLine][0:4])
      } else {
        fmt.Print(lines[(int(char) - 32) * 8 + iLine])
      }
    }
    fmt.Println()
  }
}

func Basic(str, banner string) {

  for i := 0; i < len(str); {
    endStringIndex := piscine.Index(str[i: ], "\\n")
    if endStringIndex == -1 {
      endStringIndex = len(str)
    }
    printBasic(str[i: endStringIndex], banner)
    // adding the length of the printed string
    i = endStringIndex
    for ; i + 1 < len(str) && str[i: i + 2] == "\\n"; i += 2 {
      fmt.Println()
    }
	}
}

