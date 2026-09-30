package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"golang.org/x/term"
)

type Prompt struct {
	Source string
	Text   string
}

func main() {
	prompts := []Prompt{
		// Brian Eno - Oblique Strategies
		{Source: "Brian Eno", Text: "Honor thy error as a hidden intention."},
		{Source: "Brian Eno", Text: "Use an old idea."},
		{Source: "Brian Eno", Text: "What would your closest friend do?"},
		{Source: "Brian Eno", Text: "Displace an element."},
		{Source: "Brian Eno", Text: "Remove specifics and convert to probabilities."},
		{Source: "Brian Eno", Text: "Is there something missing?"},
		{Source: "Brian Eno", Text: "Give the game to someone else."},
		{Source: "Brian Eno", Text: "Emphasize differences."},
		{Source: "Brian Eno", Text: "Look closely at the most embarrassing detail and amplify it."},
		{Source: "Brian Eno", Text: "Think - like vertical clouds, like horizontal trees."},
		{Source: "Brian Eno", Text: "Work at a different speed."},
		{Source: "Brian Eno", Text: "Into the impossible."},
		{Source: "Brian Eno", Text: "Courage!"},
		{Source: "Brian Eno", Text: "Change instrument roles."},

		// Bert Ligon - Compositional & Line Transformation Techniques
		{Source: "Bert Ligon", Text: "Reverse: Play the melodic phrase in retrograde from the end to the beginning."},
		{Source: "Bert Ligon", Text: "Change Speed: Apply rhythmic augmentation (double note values) or diminution (halve them)."},
		{Source: "Bert Ligon", Text: "Mode Change: Shift the entire line from its current parent mode into parallel minor, major, or harmonic minor."},
		{Source: "Bert Ligon", Text: "Embellish: Decorate target pitches with upper or lower chromatic neighbor tones and enclosures."},
		{Source: "Bert Ligon", Text: "Inversion: Flip the interval directions—where the original line ascends by step or leap, descend by the same interval."},
		{Source: "Bert Ligon", Text: "Rhythmic Displacement: Shift the starting point of the phrase off the beat (e.g., from beat 1 to the 'and' of 1)."},
		{Source: "Bert Ligon", Text: "Fragmentation: Take a 2-bar motif, isolate a single 3-note segment, and repeat it over shifting harmonic contexts."},
		{Source: "Bert Ligon", Text: "Interpolation: Insert passing tones or chromatic connectors inside a previously direct arpeggiated movement."},
		{Source: "Bert Ligon", Text: "Change Range/Register: Transpose specific intervals of the line up or down an octave to break up step-wise contour."},
		{Source: "Bert Ligon", Text: "Truncation: Chop off the expected resolution of the line and leave the phrase hanging open on an unaccented pitch."},
	}

	fd := int(os.Stdin.Fd())

	// Put terminal into raw mode to read single keypresses immediately
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Printf("Error setting raw mode: %v\r\n", err)
		return
	}
	defer term.Restore(fd, oldState)

	// Clear screen & hide cursor for clean presentation
	fmt.Print("\033[2J\033[H\033[?25l")
	defer fmt.Print("\033[?25h") // Restore cursor on exit

	printPrompt(prompts[rand.IntN(len(prompts))])

	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}

		char := buf[0]

		// Spacebar triggers a new pick
		if char == ' ' {
			printPrompt(prompts[rand.IntN(len(prompts))])
		}

		// Quit on 'q', 'Q', or Ctrl+C (ASCII 3)
		if char == 'q' || char == 'Q' || char == 3 {
			fmt.Print("\r\n\r\nExiting...\r\n")
			break
		}
	}
}

func printPrompt(p Prompt) {
	// Clear terminal and move cursor to top-left
	fmt.Print("\033[2J\033[H")

	fmt.Print("================================================================================\r\n")
	fmt.Print("                        COMPOSITIONAL PROMPT GENERATOR                         \r\n")
	fmt.Print("================================================================================\r\n\r\n")

	fmt.Printf(" [SOURCE]: %s\r\n\r\n", p.Source)
	fmt.Printf("   \"%s\"\r\n\r\n", p.Text)

	fmt.Print("--------------------------------------------------------------------------------\r\n")
	fmt.Print(" Press [SPACEBAR] for next prompt  |  Press [Q] to exit\r\n")
}