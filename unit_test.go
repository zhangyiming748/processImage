package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Volumes/swap/Back"
	pattern := "jpg;JPG;png;PNG;jpeg"
	threads := "4"
	ProcessAllImages(dir, pattern, threads)
}
