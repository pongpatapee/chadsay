package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const FIGURES_DIR = "./figures"

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

func readFigureFile(figureName string) (string, error) {
	figureFilePath := fmt.Sprintf("%v/%v.txt", FIGURES_DIR, figureName)

	figure, err := os.ReadFile(figureFilePath)
	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return string(figure), nil
}

func getValidFigures() []string {
	entries, err := os.ReadDir(FIGURES_DIR)
	if err != nil {
		log.Fatal(err)
	}

	var figures []string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		filename := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))

		figures = append(figures, filename)
	}

	return figures
}

func printValidFigures() {
	validFigures := getValidFigures()
	fmt.Println("Valid figure options:")
	for _, figure := range validFigures {
		fmt.Println(figure)
	}
}

func printTextBubble(balloon string) {
	bubbleLine := `    \
     \
      \
    `

	fmt.Println(balloon)
	fmt.Print(bubbleLine)
}

func printFigure(figure string) {
	figure = strings.Replace(figure, "\n", "\n\t", -1)
	fmt.Printf("\t")
	fmt.Println(figure)
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Command is intended to be used with pipes.")
		fmt.Println("Example: fortune | chadsay")
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

	var figureType string
	var listFlag bool
	flag.StringVar(&figureType, "f", "chad", "Enter the figure name. Ex: chadsay -f <figureName>. Get valid figures with chadsay -l")
	flag.BoolVar(&listFlag, "l", false, "Lists out available figures")
	flag.Parse()

	if listFlag {
		printValidFigures()
		return
	}

	lines = tabsToSpaces(lines)
	maxWidth := calculateMaxWidth(lines)
	message := normalizeStringsLength(lines, maxWidth)
	balloon := buildBalloon(message, maxWidth)
	figure, _ := readFigureFile(figureType)

	printTextBubble(balloon)
	printFigure(figure)
}
