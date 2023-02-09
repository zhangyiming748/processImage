package processImage

import (
	"bufio"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/processImage/util"
	"github.com/zhangyiming748/voiceAlert"
	"os"
	"strings"
	"time"
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
			// DeleteUnfinishedFile()
			voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "文件转换失败")
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
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
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	voiceAlert.CustomizedOnMac(voiceAlert.Shanshan, "单个目录下文件全部转换完成")
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}
func ProcessAllImages(root, pattern, threads string) {
	ProcessImages(root, pattern, threads)
	Folders := GetAllFolder.ListFolders(root)
	for _, Folder := range Folders {
		ProcessImages(Folder, pattern, threads)
	}
}
func ProcessImagesLikeGif(dir, pattern, threads string) {
	defer func() {
		if err := recover(); err != nil {
			voiceAlert.Voice(voiceAlert.FAILED)
		}
	}()
	m_start := time.Now()
	start := time.Now().Format("整个任务开始时间 15:04:03")
	log.Debug.Println(start)
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		util.Dynamic(file, threads)
		voiceAlert.Voice(voiceAlert.SUCCESS)
	}
	m_end := time.Now()
	end := time.Now().Format("整个任务结束时间 15:04:03")
	log.Debug.Println(end)
	during := m_end.Sub(m_start).Minutes()
	voiceAlert.Voice(voiceAlert.COMPLETE)
	log.Debug.Printf("整个任务用时 %v 分\n", during)
}
func DeleteUnfinishedFile() {
	defer func() {
		if err := recover(); err != nil {
			log.Warn.Println("文件已经删除过了")
		}
	}()
	f := readLastLine()
	log.Debug.Println("即将删除的文件", f)
	os.RemoveAll(f)
}

func readLastLine() (info string) {
	file, err := os.Open("program.log")
	if err != nil {
		log.Warn.Fatal(err)
	}
	defer file.Close()
	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
		//fmt.Print(lineText)
	}
	return strings.Split(string(lineText), "threads 2")[2]
}
