package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func buildBalloon(lines []string, maxWidth int) string {
	borders := []string{
		"/", "\\", "\\", "/", "|", "<", ">",
	}

	count := len(lines)

	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)

	var balloon []string
	balloon = append(balloon, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		balloon = append(balloon, s)

	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		balloon = append(balloon, s)

		i := 1
		for ; i < count-1; i++ {

			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			balloon = append(balloon, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		balloon = append(balloon, s)
	}

	balloon = append(balloon, bottom)

	return strings.Join(balloon, "\n")
}

func tabsToSpaces(lines []string) []string {
	var output []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		output = append(output, l)
	}

	return output
}

func calculateMaxWidth(lines []string) int {
	maxWidth := 0
	for _, l := range lines {
		length := utf8.RuneCountInString(l)
		if length > maxWidth {
			maxWidth = length
		}
	}

	return maxWidth
}

func normalizeStringsLength(lines []string, maxWidth int) []string {
	var newLines []string

	for _, l := range lines {
		newLine := l + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(l))
		newLines = append(newLines, newLine)
	}

	return newLines
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Command is intended to be used with pipes.")
		fmt.Println("Example: fotune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(input))
	}

	cow := `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines)
	maxWidth := calculateMaxWidth(lines)
	message := normalizeStringsLength(lines, maxWidth)
	balloon := buildBalloon(message, maxWidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
