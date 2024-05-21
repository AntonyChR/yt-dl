package internal

import (
	"os"
	"testing"
)

func TestSendToTelegramChat(t *testing.T) {
	telegram := Telegram{
		BaseUrl:  "https://api.telegram.org/bot",
		BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChatId:   os.Getenv("TELEGRAM_CHAT_ID"),
	}

	if telegram.BotToken == "" {
		t.Error("TELEGRAM_BOT_TOKEN is not set")
	}

	if telegram.ChatId == "" {
		t.Error("TELEGRAM_CHAT_ID is not set")
	}

	audioPath := "../public/mp3/test.mp3"

	err := SendAudioFileToTelegram(telegram, audioPath, "Test audio file")

	if err != nil {
		t.Errorf("Error sending audio file to telegram: %s", err)
	}
}
