package game

import (
	"fmt"
	"math/rand"
	"tay/pkg/utils"
	"time"
)

type lyricLine struct {
	line  string
	verse string
}

func PlayRandomLyricsMode(song []Verse) Result {
	result := Result{}

	// flatten song into lyrics slice
	lyrics := []lyricLine{}
	for _, v := range song {
		for _, l := range v.Lines {
			lyrics = append(lyrics, lyricLine{l, v.Name})
		}
	}

	fmt.Printf("\nHow many lines do you want to test?\n")
	n := utils.GetPositiveIntInput()
	fmt.Printf("I'll give you a random line from the song, give me the next line\n\n")

	for i := 0; i < n; i++ {
		l1, l2 := getTwoLines(lyrics)
		fmt.Printf("Question %d of %d:\nVerse: %s\n'%s'\n", i+1, n, l1.verse, l1.line)
		input := utils.GetInput()
		if input == l2.line {
			fmt.Println("✅")
			result.Correct++
		} else {
			fmt.Printf("❌ The line was '%s'\n", l2.line)
			result.Wrong++
		}

		fmt.Println()
	}

	return result
}

func getTwoLines(lines []lyricLine) (lyricLine, lyricLine) {
	numLines := len(lines)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(numLines - 1)
	return lines[num], lines[num+1]
}
