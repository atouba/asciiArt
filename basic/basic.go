package basic

import (
	"fmt"
	"os"
	"github.com/atouba/piscine"
)

func Basic(str, banner string) {
	chars, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}

	inline_strings := piscine.Split(str, "\n")
	lines := piscine.Split(string(chars), "\n")
	for i_str, inline_str := range inline_strings {
		for i_octal_line := range 8 {
			for _, char := range inline_str {
				if char == ' ' {
					fmt.Print(lines[(int(char) - 32) * 8 + i_octal_line][0:4])
				} else {
					fmt.Print(lines[(int(char) - 32) * 8 + i_octal_line])
				}
			}
			fmt.Println()
		}
		if i_str + 1 < len(inline_strings) && inline_strings[i_str + 1] == "" {
			fmt.Println()
		}
	}

	// fmt.Println(len(lines))
}
