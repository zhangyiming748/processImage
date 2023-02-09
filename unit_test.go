package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	dir := "/Volumes/T7/slacking/Telegram/latex"
	pattern := "jpg;png"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
}

func TestDeleteUnfinishedFile(t *testing.T) {
	DeleteUnfinishedFile()
}
