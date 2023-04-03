package processImage

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processImage/util"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
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

		slog.Info(fmt.Sprintf("正在处理第 %d/%d 个文件", index+1, len(files)))
		//log.Debug.Printf("文件%s压缩前大小%fMB\n", file.FullName, float64(file.Size)/MegaByte)
		util.Static(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllImages(root, pattern, threads string) {
	ProcessImages(root, pattern, threads)
	Folders := GetAllFolder.ListFolders(root)
	for index, Folder := range Folders {
		slog.Info(fmt.Sprintf("正在处理第 %d/%d 个文件夹", index+1, len(Folders)))
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
