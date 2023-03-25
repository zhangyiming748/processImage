package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Users/zen/Downloads/NecDaz"
	pattern := "jpg;JPG;png;PNG;jpeg"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}
func TestProcessImages6(t *testing.T) {
	dir := "F:\\整理\\LOL"
	pattern := "jpg;JPG;png;PNG"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}
