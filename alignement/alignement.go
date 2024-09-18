package alignement

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/atouba/piscine"
)

// getTermWidth returns the width of the terminal in characters
func getTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outSplit := strings.Split(string(out), " ")
	width, _ := strconv.Atoi(outSplit[1][:len(outSplit[1])-1])
	return width
}

// alignLCR justifies the output text left, center or right
func AlignLCR(s, a string) string {
	rows := strings.Split(s, "\n")
	nuRows := []string{}
	w := getTermWidth()

	for _, row := range rows {
		add0 := ""
		add1 := ""

		if a == "left" {
			for i := 0; i < w-len(row); i++ {
				add0 += " "
			}
			row = row + add0
		}

		if a == "center" {
			for i := 0; i < (w-len(row))/2; i++ {
				add0 += " "
				add1 += " "
			}
			if (w-len(row))%2 != 0 {
				add1 += " "
			}
			row = add0 + row + add1
		}

		if a == "right" {
			for i := 0; i < w-len(row); i++ {
				add0 += " "
			}
			row = add0 + row
		}

		nuRows = append(nuRows, row)
	}
	return strings.Join(nuRows, "\n")
}

// clearCarReturns returns the input string without carriage returns
func ClearCarReturns(s string) (out string) {
	for _, r := range s {
		if r != 13 {
			out += string(r)
		}
	}
	return
}

// ---------------------------------------------------------------------------

func AsciiArtLength(str string, f func(rune) bool, style string) int {
	asciiArtChars, err := os.ReadFile("./banners/" + style + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return 0
	}

	count := 0

	//lines := piscine.Split(string(asciiArtChars), "\n")
	toLines := ClearCarReturns(string(asciiArtChars))
	lines := piscine.Split(toLines, "\n")
	for _, char := range str {
		if int(char) < 32 || int(char) > 126 {
			log.Fatal("Bad Input")
		}
		if f(char) {
			count += len(lines[(int(char)-32)*8])
		}
	}

	return count
}

// Spaces that will be added in alignement
func SpacesCount(str, alignFlag string, style string) (int, int, int) {
	var f1 func(rune) bool = func(c rune) bool { return true }
	var f2 func(rune) bool = func(c rune) bool {
		return c != ' '
	}
	// spaces to be added after str (even if it has spaces). Used for anything but justify
	subtract := getTermWidth() - AsciiArtLength(str, f1, style)
	// total nbr of spaces left after writing the ascii art chars. Used just for justify
	subtractChars := getTermWidth() - AsciiArtLength(str, f2, style)
	if subtract < 0 || subtractChars < 0 {
		log.Fatal("characters don't fit in the terminal width size")
	}
	if alignFlag == "left" {
		return 0, subtract, 0
	} else if alignFlag == "right" {
		return subtract, 0, 0
	} else if alignFlag == "center" {
		return subtract / 2, subtract / 2, 0
	} else if alignFlag == "justify" {
		countWords := len(strings.Fields(str))
		if countWords <= 1 {
			countWords = 1
		} else {
			countWords--
		}
		return 0, 0, subtractChars / countWords
	}
	return 0, 0, 0
}

// Returns n spaces string
func SpacesString(n int) string {
	s := ""

	for ; n > 0; n-- {
		s += " "
	}
	return s
}

func skipCharLength(str string, char rune) int {
	count := 0

	for _, v := range str {
		if v != char {
			count++
		}
	}
	return count
}
