package controllers

import (
	"html/template"
	"net/http"
	"yt-dl/internal"

	"github.com/google/uuid"
)

type Controller struct {
	PublicDir        string
	MP3Dir           string
	TelegramBotToken string
	TelegramChatId   string
	LogSSEManager    *internal.LogSSEManager
}

type PageData struct {
	Title     string
	Files     []string
	PublicDir string
	FilesDir  string
}

func (c *Controller) FileList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/templates/files.html"))
	files, _ := internal.GetFiles(c.MP3Dir)
	pageData := PageData{
		Title:     "MP3 files",
		PublicDir: c.PublicDir,
		Files:     files,
		FilesDir:  c.MP3Dir,
	}
	tmpl.Execute(w, pageData)
}

func (c *Controller) SendToTelegram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log := "==>> Sending file: " + fileName + " to telegram"
	println(log)
	c.LogSSEManager.GreenLog(log)
	telegram := internal.Telegram{
		BaseUrl:  "https://api.telegram.org/bot",
		BotToken: c.TelegramBotToken,
		ChatId:   c.TelegramChatId,
	}
	internal.SendAudioFileToTelegram(telegram, c.MP3Dir+"/"+fileName, "")
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DownloadVideoAsMp3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	videoUrl := r.FormValue("url")
	fileName := r.FormValue("fileName")
	if videoUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log := "==>> Downloading video from: " + videoUrl
	println(log)
	c.LogSSEManager.GreenLog(log)
	err := internal.DownloadVideo(videoUrl, c.MP3Dir, fileName, c.LogSSEManager)
	if err != nil {
		c.LogSSEManager.RedLog(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DeleteFileByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log := "==>> Deleting file: " + fileName
	println(log)
	c.LogSSEManager.RedLog(log)
	err := internal.DeleteFile(c.MP3Dir + "/" + fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Controller) ServerLogsSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientId := uuid.New().String()

	c.LogSSEManager.Register(w, clientId)

	println("==>> Client connected:", clientId)
	w.(http.Flusher).Flush()

	done := r.Context().Done()
	for {
		<-done
		println("==>> Client disconnected: ", clientId)
		c.LogSSEManager.Unregister(clientId)
		return
	}
}
