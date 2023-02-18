package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Volumes/T7/slacking/Telegram"
	pattern := "jpg;JPG;png;PNG"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}
