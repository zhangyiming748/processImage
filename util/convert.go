package util

import (
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

func Static(in File, threads string) {
	prefix := strings.Trim(in.FullPath, in.ExtName)
	log.Debug.Printf("prefix is %v\n", prefix)
	out := strings.Join([]string{prefix, "avif"}, ".")
	log.Debug.Printf("in is %v\n", in.FullPath)
	log.Debug.Printf("out is %v\n", out)
	// ffmpeg -i image.png -c:v libaom-av1 -still-picture 1 image.avif
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libaom-av1", "-still-picture", "1", "-threads", threads, out)
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = replace.Replace(t)
		log.Info.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Panicf("命令执行中有错误产生:%v\n", err)
	}
	//log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, file)
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Debug.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
}
func Dynamic() {
	// TODO ffmpeg -i video.mp4 animated.avif
}
