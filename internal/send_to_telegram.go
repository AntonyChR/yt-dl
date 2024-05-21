package internal

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

//var TELEGRAM_BASE_URL = "https://api.telegram.org/bot" + TELEGRAM_BOT_TOKEN

type Telegram struct {
	BaseUrl  string
	BotToken string
	ChatId   string
}

func SendAudioFileToTelegram(telegram Telegram, relativePath, caption string) error {
	absPath := relativePath
	if strings.Contains(relativePath, "~/") {
		userHomeDir, _ := os.UserHomeDir()
		absPath = strings.ReplaceAll(relativePath, "~/", userHomeDir+"/")
	}

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
		return err
	}
	defer file.Close()
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	part, _ := writer.CreateFormFile("audio", path.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("chat_id", telegram.ChatId)
	writer.WriteField("caption", caption)
	writer.Close()

	req, _ := http.NewRequest("POST", telegram.BaseUrl+telegram.BotToken+"/sendAudio", &body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {

		fmt.Printf("Error opening file: %s", err)
	}
	print("==>> Audio file sent to telegram")
	return err
}
