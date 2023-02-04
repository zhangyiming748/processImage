package processImage

import "testing"

func TestProcessImages(t *testing.T) {
	dir := "/Volumes/T7/slacking/Telegram/nier"
	pattern := "jpg;png"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}
