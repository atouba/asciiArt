package alignement

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atouba/piscine"
	"golang.org/x/term"
)

// getTermWidth returns the width of the terminal in characters
func getTermWidth() int {
	/* 	cmd := exec.Command("stty", "size")
	   	cmd.Stdin = os.Stdin
	   	out, err := cmd.Output()

	   	if err != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	}

	   	outSplit := strings.Split(string(out), " ")
	   	width, _ := strconv.Atoi(outSplit[1][:len(outSplit[1])-1]) */

	width, _, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return width
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

// AsciiArtLength returns the length of a line of a ascii art from the string str
func AsciiArtLength(str string, f func(rune) bool, style string) int {
	asciiArtChars, err := os.ReadFile("./banners/" + style + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return 0
	}

	count := 0

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

// SpacesCount tells how many spaces should be added in different places to the output for alignement
func SpacesCount(str, alignFlag string, style string) (int, int, int) {

	if alignFlag == "" {
		return 0, 0, 0
	}

	var f1 func(rune) bool = func(c rune) bool { return true }
	var f2 func(rune) bool = func(c rune) bool {
		return c != ' '
	}
	// spaces to be added after str (even if it has spaces). Used for anything but justify
	subtract := getTermWidth() - AsciiArtLength(str, f1, style)
	// total nbr of spaces left after writing the ascii art chars. Used just for justify
	subtractChars := getTermWidth() - AsciiArtLength(str, f2, style)
	if (alignFlag != "justify" && subtract < 0) || (alignFlag == "justify" && subtractChars < 0) {
		log.Fatal("characters don't fit in the terminal width size")
	}
	if alignFlag == "left" {
		return 0, subtract, 0
	} else if alignFlag == "right" {
		return subtract, 0, 0
	} else if alignFlag == "center" {
		if subtract%2 == 0 {
			subtract--
		}
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

// SpacesString returns a string of n spaces
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
