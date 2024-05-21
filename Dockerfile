FROM golang:1.22.3-alpine3.19
WORKDIR /app
COPY . .
RUN apk add --no-cache yt-dlp
RUN apk add --no-cache ffmpeg
RUN go build -o app main.go 
RUN chmod +x start.sh
CMD ["./start.sh"]

