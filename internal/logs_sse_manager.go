package internal

import (
	"fmt"
	"net/http"
)

// NewLogSSEManager creates a new instance of LogSSEManager.
func NewLogSSEManager() *LogSSEManager {
	return &LogSSEManager{
		Clients:     make(map[string]http.ResponseWriter),
		LogsChannel: make(chan string),
	}
}

// LogSSEManager manages Server-Sent Events (SSE) for logging purposes.
type LogSSEManager struct {
	Clients     map[string]http.ResponseWriter
	LogsChannel chan string
}

func (l *LogSSEManager) Register(w http.ResponseWriter, clientId string) {
	l.Clients[clientId] = w
}

func (l *LogSSEManager) Unregister(clientId string) {
	delete(l.Clients, clientId)
}

// Start starts the LogSSEManager and listens for logs on the LogsChannel.
func (l *LogSSEManager) Start() {
	for {
		log := <-l.LogsChannel
		l.Broadcast(log)
	}
}

// Broadcast sends the log content to all registered clients.
func (l *LogSSEManager) Broadcast(content string) {
	if len(l.Clients) == 0 {
		return
	}
	data := fmt.Sprintf("<p>%s</p>", content)

	for _, client := range l.Clients {
		fmt.Fprintf(client, "data: %s\n\n", data)
		client.(http.Flusher).Flush()
	}
}

func (l *LogSSEManager) Log(log string) {
	l.LogsChannel <- "<p>" + log + "</p>"
}

func (l *LogSSEManager) RedLog(log string) {
	l.LogsChannel <- "<p class=\"text-red-600\">" + log + "</p>"
}

func (l *LogSSEManager) GreenLog(log string) {
	l.LogsChannel <- "<p class=\"text-green-600\">" + log + "</p>"
}
