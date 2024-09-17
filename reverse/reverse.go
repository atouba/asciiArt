package reverse

import (
	"fmt"
	"log"
	"os"

	"github.com/atouba/piscine"
)

func index(str string, char rune, start_index int) int {
	i := start_index
	for _, v := range str {
		if v == char {
			return i
		}
		i++
	}
	return i
}

func isVerticallyEqual(lines []string, column_index int) bool {
	lines_n := len(lines) - 1 // minus 1 so when iterating we don't avoid out of range error

	for row := 0; row < lines_n; row++ {
		if lines[row][column_index] != lines[row + 1][column_index] {
			return false
		}
	}
	return true
}

func isSpace(lines []string, column_index int) bool {
	char := lines[0][column_index]

	if char == ' ' && isVerticallyEqual(lines, column_index) {
		return true
	}
	return false
}

func isEndNewLine(asciiArt string) bool {
  // no need to check the last char as well, becasue it'll always be a new line
  return asciiArt[len(asciiArt) - 2] == '\n'
}

func asciiArtCharLength(lines []string, start_index int) int {
	length := 0

  // check first if it's an ascii art space
  x := start_index
  for ; x < start_index + 6 ; x++ {
    if !(x < len(lines[0])) { log.Fatal("Bad Input") }
    if !isSpace(lines, x) { break }
  }
  if x == start_index + 6 { return 6 }

	for i := start_index; i < len(lines[0]); i++ {
		if !isSpace(lines, i) {
			length++
		} else {
			break
		}
	}
	return length + 1
}

func Reverse(filename string) {
	ascii_art, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	lines := piscine.Split(string(ascii_art), "\n")
  lines = lines[: len(lines)-1]

  for startLines := 0; startLines < len(lines); startLines += 8 {
    for column_i := 0; column_i < len(lines[startLines: startLines+8][0]); column_i += asciiArtCharLength(lines[startLines: startLines+8], column_i) {
      printRegularChar(lines[startLines: startLines+8], column_i)
    }
    if startLines + 8 < len(lines) { fmt.Println() }
  }
  if isEndNewLine(string(ascii_art)) { fmt.Println() }
}

func printRegularChar(lines []string, column_index int) {
	bannerChars, err := os.ReadFile("./banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading standard.txt")
    return
	}

	bannerLines := piscine.Split(string(bannerChars), "\n")
	var i int // index of input line
	charLength := asciiArtCharLength(lines, column_index)

	for bannerLIndex := 0; bannerLIndex < len(bannerLines) / 8; bannerLIndex++ {
		tempIndex := bannerLIndex * 8
		for i = 0; i < 8; i++ {
			if bannerLines[tempIndex] != lines[i][column_index: column_index + charLength] {
				break
			}
			tempIndex++
		}
		if i == 8 {
			fmt.Print(string(rune(bannerLIndex + 32)))
			return
		}
	}
}
