package main

import (
	"net/http"
	"os"
	"yt-dl/internal"
	"yt-dl/internal/controllers"
)

var PUBLIC_DIR = "public"
var MP3_DIR = PUBLIC_DIR + "/mp3"

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
var TELEGRAM_CHAT_ID = os.Getenv("TELEGRAM_CHAT_ID")

func main() {

	err := internal.CheckRequiredDirectories(PUBLIC_DIR, MP3_DIR)

	if err != nil {
		panic(err)
	}

	logSSEManager := internal.NewLogSSEManager()
	go logSSEManager.Start()

	controller := controllers.Controller{
		PublicDir:        PUBLIC_DIR,
		MP3Dir:           MP3_DIR,
		TelegramBotToken: TELEGRAM_BOT_TOKEN,
		TelegramChatId:   TELEGRAM_CHAT_ID,
		LogSSEManager:    logSSEManager,
	}

	http.HandleFunc("POST /dl", controller.DownloadVideoAsMp3)

	// get "file" from query string and send it to telegram
	http.HandleFunc("POST /send", controller.SendToTelegram)

	http.HandleFunc("POST /delete", controller.DeleteFileByName)

	http.HandleFunc("GET /files", controller.FileList)

	http.HandleFunc("GET /logs", controller.ServerLogsSSE)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/", fs)

	println("Server running on port http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
