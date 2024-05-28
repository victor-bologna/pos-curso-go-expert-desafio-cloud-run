package webserver

import (
	"net/http"
	"strings"
)

type WebServer struct {
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (ws *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	ws.Handlers[strings.ToUpper(method)+" "+path] = handler
}

func (ws *WebServer) Start() {
	mux := http.NewServeMux()
	for path, handler := range ws.Handlers {
		mux.HandleFunc(path, handler)
	}
	http.ListenAndServe(":"+ws.WebServerPort, mux)
}
