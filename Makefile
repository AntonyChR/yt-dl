export TELEGRAM_BOT_TOKEN := $(shell cat .env | grep TELEGRAM_BOT_TOKEN | cut -d '=' -f2)
export TELEGRAM_CHAT_ID := $(shell cat .env | grep TELEGRAM_CHAT_ID | cut -d '=' -f2)

run:
	go run main.go
build:
	go build -o app main.go

build-image:
	docker build -t yt-dl .
