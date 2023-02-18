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
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "文件转换失败")
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	if len(files) == 0 {
		voiceAlert.CustomizedOnMac(voiceAlert.Victoria, "跳过空文件夹")
	}
	for index, file := range files {
		log.Debug.Printf("正在处理第 %d/%d 个文件\n", index+1, len(files))
		log.Debug.Printf("文件%s压缩前大小%fMB\n", file.FullName, float64(file.Size)/MegaByte)
		//if file.Size < Limit {
		//	log.Debug.Printf("文件%v很小(%f),跳过压缩\n", file.FullPath, float64(file.Size))
		//	continue
		//}
		out, _ := os.Stat(util.Static(file, threads))
		resize := out.Size()
		log.Debug.Printf("文件%s压缩后大小%fMB\n", file.FullName, float64(resize)/MegaByte)
		voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个文件转换完成")
	}
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个目录下文件全部转换完成")
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
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "文件转换失败")
		}
	}()
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		util.Dynamic(file, threads)
		voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个文件转换完成")
	}
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个目录下文件全部转换完成")
}
