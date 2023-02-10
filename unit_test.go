package processImage

import (
	"os"
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Users/zen/Downloads/整理/Aiss/P"
	pattern := "jpg;JPG;png;PNG"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
	os.Exit(0)
}
