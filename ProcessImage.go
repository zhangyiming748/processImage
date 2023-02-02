package processImage

import (
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processImage/util"
	"github.com/zhangyiming748/voiceAlert"
	"time"
)

func ProcessImages(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Voice(voiceAlert.FAILED)
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	files := util.GetMultiFiles(dir, pattern)
	for _, file := range files {
		//frame := util.DetectFrame(file)
		//go getInfo.GetVideoFrame(file.FullPath)
		//convert.Convert2H265(file, threads)
		util.Static(file, threads)
		voiceAlert.Voice(voiceAlert.SUCCESS)
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	voiceAlert.Voice(voiceAlert.COMPLETE)
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}
func ProcessAllImages(root, pattern, threads string) {

}
func ProcessImagesLikeGif(dir, pattern, threads string) {

}
