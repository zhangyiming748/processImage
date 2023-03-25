package processImage

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/processImage/util"
	"github.com/zhangyiming748/voiceAlert"
	"golang.org/x/exp/slog"
	"io"
	"os"
)

func init() {
	logLevel := os.Getenv("LEVEL")
	//var level slog.Level
	var opt slog.HandlerOptions
	switch logLevel {
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
	//defer logf.Close() //如果不关闭可能造成内存泄露
	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(logger)
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
