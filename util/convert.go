package util

import (
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

/*
返回输出文件全路径
*/
func Static(in GetFileInfo.Info, threads string) string {
	prefix := strings.Trim(in.FullPath, in.ExtName)
	out := strings.Join([]string{prefix, "avif"}, ".")
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
		t := string(tmp)
		t = replace.Replace(t)
		log.Info.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
	return out
}
func Dynamic(in GetFileInfo.Info, threads string) {
	prefix := strings.Trim(in.FullPath, in.ExtName)
	out := strings.Join([]string{prefix, "avif"}, ".")
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-threads", threads, out)
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
		t := string(tmp)
		t = replace.Replace(t)
		log.Info.Printf("%v\b", t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Warn.Panicf("命令执行中有错误产生:%v\n", err)
	}
	if err := os.RemoveAll(in.FullPath); err != nil {
		log.Warn.Printf("删除源文件失败:%v\n", err)
	} else {
		log.Debug.Printf("删除源文件:%v\n", in.FullName)
	}
}
