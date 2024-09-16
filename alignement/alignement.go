package alignement

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
