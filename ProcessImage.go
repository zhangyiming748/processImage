package processImage

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processImage/util"
	"github.com/zhangyiming748/voiceAlert"
	"os"
)

const (
	Byte        = iota + 1
	KiloByte    = 1000 * Byte
	MegaByte    = 1000 * KiloByte
	Gigabyte    = 1000 * MegaByte
	Terabyte    = 1000 * Gigabyte
	Petabyte    = 1000 * Terabyte
	Exabyte     = 1000 * Petabyte
	Zettabyte   = 1000 * Exabyte
	Yottabyte   = 1000 * Zettabyte
	Brontobytes = 1000 * Yottabyte
	Geopbyte    = 1000 * Brontobytes
)

const (
	Limit = 2 * MegaByte
)

func ProcessImages(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	if len(files) == 0 {
		voiceAlert.Customize("skip", voiceAlert.Samantha)
		return
	}
	for index, file := range files {
		log.Debug.Printf("正在处理第 %d/%d 个文件\n", index+1, len(files))
		log.Debug.Printf("文件%s压缩前大小%fMB\n", file.FullName, float64(file.Size)/MegaByte)
		out, _ := os.Stat(util.Static(file, threads))
		resize := out.Size()
		log.Debug.Printf("文件%s压缩后大小%fMB\n", file.FullName, float64(resize)/MegaByte)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllImages(root, pattern, threads string) {
	ProcessImages(root, pattern, threads)
	Folders := GetAllFolder.ListFolders(root)
	for index, Folder := range Folders {
		log.Debug.Printf("正在处理第 %d/%d 个文件夹\n", index+1, len(Folders))
		ProcessImages(Folder, pattern, threads)
	}
}

func ProcessImagesLikeGif(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		util.Dynamic(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}
