package internal

import (
	"bufio"
	"fmt"
	"os/exec"
)

func DownloadVideo(url, dir, fileName string, logsChann *LogSSEManager) error {
	finalFileName := " -o \"%(uploader)-s%(title)s.%(ext)s\" "
	if fileName != "" {
		finalFileName = fmt.Sprintf(" -o \"%s.%%(ext)s\" ", fileName)
	}
	command := "yt-dlp -f bestaudio -x --audio-format mp3" + " -P " + dir + finalFileName + url
	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logsChann.RedLog("==>> Error downloading [error reading stdout pipe]: ")
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logsChann.RedLog("==>> Error downloading [error reading stderr pipe]: " + url)
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			logsChann.Log(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			logsChann.Log(scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		logsChann.RedLog("==>> Error downloading: " + url)
		return err
	}

	if err := cmd.Wait(); err != nil {
		logsChann.RedLog("==>> Error downloading (wait): " + url)
		return err
	}
	log := "==>> Downloaded video"
	println(log)
	logsChann.GreenLog(log)
	return nil
}
