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
func TestProcessImages6(t *testing.T) {
	dir := "/Volumes/Gloway720/recovery/照片"
	pattern := "jpg;JPG;png;PNG"
	threads := "10"
	ProcessAllImages6(dir, pattern, threads)
}
