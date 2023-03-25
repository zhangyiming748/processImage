package util

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"strings"
)

/*
 */
func Static(in GetFileInfo.Info, threads string) {
	prefix := strings.Trim(in.FullPath, in.ExtName)
	out := strings.Join([]string{prefix, "avif"}, ".")
	// Todo 如果转换出现错误 记录最后一个出问题的文件
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("最后出问题的文件", slog.Any("源文件", in.FullPath, ), slog.Any("目标文件", out), slog.Any("删除命令", fmt.Sprintf("rm \"%v\"", out)))
		}
	}()
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libaom-av1", "-still-picture", "1", "-threads", threads, out)
	slog.Info("ffmpeg", slog.Any("生成的命令", cmd))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		slog.Debug(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("cmd.Wait", slog.Any("错误", err))
		return
	}
	if err := os.RemoveAll(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", in.FullPath))
	}
}

/*
返回输出文件全路径
*/
func Dynamic(in GetFileInfo.Info, threads string) {
	prefix := strings.Trim(in.FullPath, in.ExtName)
	out := strings.Join([]string{prefix, "avif"}, ".")
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-threads", threads, out)
	// Todo 如果转换出现错误 记录最后一个出问题的文件
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("最后出问题的文件", slog.Any("源文件", in.FullPath, ), slog.Any("目标文件", out), slog.Any("删除命令", fmt.Sprintf("rm \"%v\"", out)))
		}
	}()
	slog.Info("ffmpeg", slog.Any("生成的命令", cmd))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		slog.Warn("cmd.StdoutPipe", slog.Any("错误", err))
		return
	}
	if err = cmd.Start(); err != nil {
		slog.Warn("cmd.Run", slog.Any("错误", err))
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = replace.Replace(t)
		slog.Debug(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		slog.Warn("cmd.Wait", slog.Any("错误", err))
		return
	}
	if err := os.RemoveAll(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
	} else {
		slog.Info("删除成功", slog.Any("源文件", in.FullPath))
	}
}
