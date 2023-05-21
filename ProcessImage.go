package processImage

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
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
		Static(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllImages(root, pattern, threads string) {
	ProcessImages(root, pattern, threads)
	Folders := GetAllFolder.List(root)
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
		Dynamic(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}
func ProcessAllImagesLikeGif(root, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	folders := GetAllFolder.List(root)
	folders = append(folders, root)
	for _, folder := range folders {
		files := GetFileInfo.GetAllFileInfo(folder, pattern)
		for _, file := range files {
			Dynamic(file, threads)
			voiceAlert.Customize("done", voiceAlert.Samantha)
		}
		voiceAlert.Customize("complete", voiceAlert.Samantha)
	}
	voiceAlert.Customize("All complete", voiceAlert.Samantha)
}
func processAllToGIF(root, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Customize("failed", voiceAlert.Samantha)
		}
	}()
	folders := GetAllFolder.List(root)
	folders = append(folders, root)
	for _, folder := range folders {
		files := GetFileInfo.GetAllFileInfo(folder, pattern)
		for _, file := range files {
			ToGIF(file, threads)
			voiceAlert.Customize("done", voiceAlert.Samantha)
		}
		voiceAlert.Customize("complete", voiceAlert.Samantha)
	}
	voiceAlert.Customize("All complete", voiceAlert.Samantha)
}
