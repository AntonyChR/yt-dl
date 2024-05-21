package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func DownloadVideo(url, dir, fileName string) error {
	finalFileName := " -o \"%(uploader)-s%(title)s.%(ext)s\" "
	if fileName != "" {
		finalFileName = fmt.Sprintf(" -o \"%s.%%(ext)s\" ", fileName)
	}
	command := "yt-dlp -f bestaudio -x --audio-format mp3" + " -P " + dir + finalFileName + url
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	println("==>> Downloaded video")
	return nil
}
