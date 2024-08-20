package asciiart

import (
	"errors"
	"os"
)

// ReadBannerFile reads the content of a banner file specified by the filepath argument and returns it as a string.
func ReadBannerFile(filepath string) (string, error) {
	bannerFile, err := os.ReadFile(filepath)

	if len(bannerFile) != 6623 && len(bannerFile) != 5558 && len(bannerFile) != 7463 && len(bannerFile) != 6262 {
		return "404", errors.New("the bannerfile has been tampered with")
	}

	if err != nil {
		return "404", err
	}
	return string(bannerFile), nil
}
