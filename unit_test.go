package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Users/zen/Pictures/emoji"
	pattern := "jpg;JPG;png;PNG"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}
