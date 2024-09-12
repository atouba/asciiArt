package reverse

import (
	"fmt"
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

func isVerticallyEqual(lines *[]string, column_index int) bool {
	lines_n := len(*lines) - 1 // minus 1 so when iterating we don't reach the last line

	for row := 0; row < lines_n; row++ {
		if (*lines)[row][column_index] != (*lines)[row + 1][column_index] {
			return false
		}
	}
	return true
}

func isSpace(lines *[]string, column_index int) bool {
	char := (*lines)[0][column_index]

	if char == ' ' && isVerticallyEqual(lines, column_index) {
		return true
	}
	return false
}

// should handle the case when there are spaces
func asciiArtCharLength(lines *[]string, start_index int) int {
	length := 0

	for i := start_index; i < len((*lines)[0]); i++ {
		if !isSpace(lines, i) {
			length++
		} else {
			break
		}
	}
	return length
}

// always assuming inline strings
func Reverse(filename string) {
	if len(filename) <= 10 || piscine.Compare(filename[0:10], "--reverse=") != 0 {
		fmt.Println(`Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>`)
		return
	}
	ascii_art, err := os.ReadFile(filename[10:])
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	lines := piscine.Split(string(ascii_art), "\n")
	
	// column_i < len(lines[0]) - 1 becasue the line has an extra char '\n'
	for column_i := 0; column_i < len(lines[0]) - 1; column_i += asciiArtCharLength(&lines, column_i) + 1 {
		// presumably lines always consist of exactly 8 lines
		printRegularChar(&lines, column_i)
	}
	fmt.Println()
}

func printRegularChar(lines *[]string, column_index int) {
	bannerChars, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading standard.txt")
	}

	bannerLines := piscine.Split(string(bannerChars), "\n")
	var i int // index of input line
	charLength := asciiArtCharLength(lines, column_index) + 1

	for bannerLIndex := 0; bannerLIndex < len(bannerLines) / 8; bannerLIndex++ {
		tempIndex := bannerLIndex * 8
		for i = 0; i < 8; i++ {
			if bannerLines[tempIndex] != (*lines)[i][column_index: column_index + charLength] {
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