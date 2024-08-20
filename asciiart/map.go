package asciiart

import (
	"strings"
)

// MapCreator creates a map of ASCII art from a string
func MapCreator(s string) map[rune][]string {
	Map := make(map[rune][]string)
	var lines []string
	printableRune := rune(32)

	// Check for how lines are split in the banner file
	if strings.ContainsRune(s, '\r') {
		lines = strings.Split(s, "\r\n")
	} else {
		lines = strings.Split(s, "\n")
	}
	for i := 0; i < len(lines)-1; i++ {
		// If the current line is empty and there are lines left to process
		if lines[i] == "" {
			artLines := []string{}
			for j := 0; j < 8; j++ {
				artLines = append(artLines, lines[i+1+j])
			}
			// Map the printable rune to its corresponding ASCII art
			Map[printableRune] = artLines
			printableRune++
		}
	}
	return Map
}
