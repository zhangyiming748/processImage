package processImage

import (
	"testing"
)

func TestProcessImages(t *testing.T) {
	setLog("Debug")
	dir := "/Users/zen/Downloads/Tele"
	pattern := "jpg;JPG;png;PNG;jpeg;bmp"
	threads := "10"
	ProcessAllImages(dir, pattern, threads)
	mylog.Debug("done!")
}

func TestProcessAllImagesLikeGif(t *testing.T) {
	pattern := "gif"
	root := "/Volumes/T7/slacking/Telegram"
	threads := "8"
	ProcessAllImagesLikeGif(root, pattern, threads)
}

//func setLog() {
//	//var level slog.Level
//	var opt slog.HandlerOptions
//	opt = slog.HandlerOptions{ // 自定义option
//		AddSource: true,
//		Level:     slog.LevelDebug, // slog 默认日志级别是 info
//	}
//	file := "test.log"
//	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
//	if err != nil {
//		panic(err)
//	}
//	//defer logf.Close() //如果不关闭可能造成内存泄露
//	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
//	slog.SetDefault(logger)
//}
