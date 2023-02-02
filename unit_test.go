package processImage

import "testing"

func TestProcessImages(t *testing.T) {
	dir := "/Users/zen/Downloads/Telegram Desktop/OverWatch"
	pattern := "jpg;png"
	threads := "4"
	ProcessImages(dir, pattern, threads)
}
