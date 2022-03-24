package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	expression    = "[^a-zA-Z0-9 ]+"
	lineSeparator = "======================" // 20 of these "=" jawns
)

func PrintSeperator() {
	fmt.Println(lineSeparator)
}

func GetEmoji(score float32) string {
	score = score * 100
	switch {
	case score >= 100:
		return "🤩"
	case score >= 90:
		return "😎"
	case score >= 80:
		return "🥸"
	case score >= 70:
		return "🤨"
	default:
		return "😔"
	}
}

func CleanText(s string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile(expression)
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	processedString = strings.ToLower(processedString)

	return processedString
}

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	line := scanner.Text()

	return CleanText(line)
}

func GetPositiveIntInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	line := scanner.Text()
	line = CleanText(line)

	i, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("Give a number!]n")
		return GetPositiveIntInput()
	} else if i <= 0 {
		fmt.Printf("Give a number greater than 0, silly!\n")
		return GetPositiveIntInput()
	}
	return i
}
