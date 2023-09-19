package asciiartweb

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CharConvert(arr []string, char rune, line int) string {
	Ascii := (int(char)-32)*9 + line
	if Ascii >= 0 && Ascii < len(arr) {
		return arr[Ascii]
	}
	return ""
}
func AsciiTest() {
	banner := "standard"
	text := ""
	hello := ""
	arg := os.Args[1:]
	if len(arg) == 3 && strings.HasPrefix(arg[0], "--output=") {
		text = arg[0][9:]
		banner = arg[2]
		hello = arg[1]
	} else if len(arg) == 1 && !strings.HasPrefix(arg[0], "--output") {
		hello = arg[0]
	} else if len(arg) == 2 {
		hello = arg[0]
		banner = arg[1]
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}
	oldStdout := os.Stdout
	if len(arg) == 3 {
		file, err := os.Create(text)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()
		// Redirect standard output to the file
		os.Stdout = file
	}
	if len(arg) == 3 {
		switch arg[2] {
		case "shadow":
			banner = "shadow"
		case "thinkertoy":
			banner = "thinkertoy"
		case "standard":
			banner = "standard"
		default:
			banner = "standard"
		}
	}
	bannerContent, err := ioutil.ReadFile(fmt.Sprintf("%s.txt", banner))
	if err != nil {
		fmt.Println("error reading file: ", err)
		os.Exit(1)
	}
	lines := strings.Split(string(bannerContent), "\n")
	// Split the input string into lines using the "\n" escape sequence
	splitted := strings.Split(hello, `\n`)
	// Loop through each line of the input string and each row of the banner
	num := 0
	// Loop through each line of the input string and each row of the banner
	for _, line := range splitted {
		num = num + 1
		if line == "" {
			// Print a newline if the input line is empty
			if num < len(splitted) {
				fmt.Println()
				continue
			} else {
				continue
			}
		} else {
			for j := 1; j < 9; j++ {
				// Convert each character of the input line to the corresponding ASCII art character in the banner
				for _, char := range line {
					fmt.Print(CharConvert(lines, rune(char), j))
				}
				fmt.Println()
			}
		}
	}
	os.Stdout = oldStdout
	// Restore standard output
}
