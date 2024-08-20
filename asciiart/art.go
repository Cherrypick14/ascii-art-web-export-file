package asciiart

// reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the console.
func Art(inputText, banner string) (string, error) {
	bannerFile, err := ReadBannerFile("banners/" + banner + ".txt")
	if err != nil {
		return "404", err
	}
	ASCIIArtMap := MapCreator(string(bannerFile))

	asciiArt, err := ArtRetriever(inputText, ASCIIArtMap)
	if err != nil {
		return "400", err
	}
	return asciiArt, nil
}
