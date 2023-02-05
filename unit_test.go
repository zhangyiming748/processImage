package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Users/zen/Downloads/Affect3d/futa"
	pattern := "jpg;png"
	threads := "2"
	ProcessAllImages(dir, pattern, threads)
}

func TestDeleteUnfinishedFile(t *testing.T) {
	DeleteUnfinishedFile()
}
