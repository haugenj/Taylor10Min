package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tay/pkg/game"
	"tay/pkg/utils"
)

// Options configurable via flags or environment variables
type Options struct {
	verse    bool
	numLines int
}

var options = Options{}
var verses = []game.Verse{}

//go:embed data
var f embed.FS

func main() {
	initSigs()

	flag.BoolVar(&options.verse, "verse", utils.WithDefaultBool("VERSE", false), "Test whole verses instead of lines")
	flag.IntVar(&options.numLines, "numLines", utils.WithDefaultInt("NUM_LINES", 5), "How many lines you want to test")
	flag.Parse()

	// Parse the lyrics file

	file, err := f.Open("data/verses.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	verseName := ""
	var verse game.Verse
	for scanner.Scan() {
		line := scanner.Text()
		// check if it's a title or a line
		if strings.Index(line, "[") == 0 {
			if verseName != "" {
				verses = append(verses, verse)
			}
			// new verse
			verseName = line
			verse.Name = verseName
			verse.Lines = []string{}
		} else {
			verse.Lines = append(verse.Lines, utils.CleanText(line))
		}
	}

	// Prompt for game type

	utils.PrintSeperator()
	fmt.Printf("Welcome to the ✨ All Too Well (10 min Taylor's Version) study guide ✨\n")
	utils.PrintSeperator()
	fmt.Println()

	fmt.Printf("Select a game mode:\n")
	fmt.Printf("1 - Random Lines Test\n")
	fmt.Printf("2 - Full Verse Test\n")

	// Run the game
	result := playGame(verses)

	// Print Results

	printResults(result)
}

func initSigs() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		// maybe handle this better
		printResults(game.Result{})
		os.Exit(0)
	}()
}

func printResults(result game.Result) {
	fmt.Println()
	utils.PrintSeperator()

	total := result.Correct + result.Wrong
	if total == 0 {
		fmt.Printf("You didn't even try!\n")
	} else {
		score := float32(result.Correct) / float32(total)
		fmt.Printf("You got %d out of %d correct for a final score of %.2f %s\n", result.Correct, total, score*100, utils.GetEmoji(score))
	}

	utils.PrintSeperator()
}

func playGame(song []game.Verse) game.Result {
	input := utils.GetInput()
	switch {
	case input == "1":
		return game.PlayRandomLyricsMode(song)
	case input == "2":
		return game.PlayVerseMode(song)
	default:
		fmt.Printf("Enter 1, or 2 >:(\n")
		return playGame(song)
	}
}
