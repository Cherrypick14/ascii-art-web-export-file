package asciiart

import (
	"fmt"
	"strings"
)

// ArtRetriever returns the ASCII art corresponding to the input string using the provided map.
func ArtRetriever(s string, m map[rune][]string) (string, error) {
	var result strings.Builder

	if s == "" {
		result.WriteString("")
		return result.String(), nil
	}
	lines := strings.Split(s, "\n")

	for ind := 0; ind < len(lines); ind++ {
		if lines[ind] == "" {
			result.WriteString("\n")
		} else {
			for j := 0; j < 8; j++ {
				for _, char := range lines[ind] {
					if asciiArt, ok := m[char]; ok {
						// Add the corresponding ASCII art for the character
						result.WriteString(asciiArt[j])
					} else {
						// For cases of non-ascii characters
						return "400", fmt.Errorf("error! invalid input: %s", string(char))
					}
				}
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}
