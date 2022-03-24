package game

import (
	"fmt"
	"math/rand"
	"tay/pkg/utils"
	"time"
)

func PlayVerseMode(song []Verse) Result {
	result := Result{}

	fmt.Printf("How many verses do you want to test?\n")
	n := utils.GetPositiveIntInput()
	fmt.Printf("I'll give you a verse and the first line, enter the rest of the verse line by line\n\n")

	numVerses := len(song)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		v := song[rand.Intn(numVerses)]

		fmt.Printf("Question %d of %d:\nVerse: %s\n", i+1, n, v.Name)
		fmt.Printf("'%s'\n", v.Lines[0])
		for _, l := range v.Lines[1:] {
			input := utils.GetInput()
			if input == l {
				fmt.Println("✅")
				result.Correct++
			} else {
				fmt.Printf("❌ The line was '%s'\n", l)
				result.Wrong++
			}
		}

		fmt.Println()
	}

	return result
}
