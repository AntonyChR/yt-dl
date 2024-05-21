package controllers

import (
	"html/template"
	"net/http"
	"yt-dl/internal"
)

type Controller struct {
	PublicDir        string
	MP3Dir           string
	TelegramBotToken string
	TelegramChatId   string
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
	println("==>> Sending file: " + fileName + " to telegram")
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
	println("==>> Downloading video from: " + videoUrl)
	err := internal.DownloadVideo(videoUrl, c.MP3Dir, fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DeleteFileByFileName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	println("==>> Deleting file: " + fileName)
	err := internal.DeleteFile(c.MP3Dir + "/" + fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
