package processImage

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"io"
	"os"
)

var mylog *slog.Logger

func setLog(level string) {
	var opt slog.HandlerOptions
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	file := "processImage.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	mylog = slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
}
func init() {
	l := os.Getenv("LEVEL")
	setLog(l)
}
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
		mylog.Info(fmt.Sprintf("正在处理第 %d/%d 个文件", index+1, len(files)))
		Static(file, threads)
		voiceAlert.Customize("done", voiceAlert.Samantha)
	}
	voiceAlert.Customize("complete", voiceAlert.Samantha)
}

func ProcessAllImages(root, pattern, threads string) {
	ProcessImages(root, pattern, threads)
	Folders := GetAllFolder.List(root)
	for index, Folder := range Folders {
		mylog.Info(fmt.Sprintf("正在处理第 %d/%d 个文件夹", index+1, len(Folders)))
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
