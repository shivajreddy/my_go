package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var (
	counter int
	clients = make(map[chan int]bool)
	users   = make(map[string]bool)
)

func main() {
	fmt.Println("LEARNING REALTIME DEV")
	http.HandleFunc("/stream", stream)
	http.HandleFunc("/", home)
	go tick()
	fmt.Println("server on :6969")
	http.ListenAndServe(":6969", nil)
}

func tick() {
	for {
		time.Sleep(time.Second)
		counter++
		for ch := range clients {
			select {
			case ch <- counter:
			default:
			}
		}
	}
}

func stream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	ch := make(chan int)
	clients[ch] = true
	defer func() { delete(clients, ch); close(ch) }()
	for count := range ch {
		fmt.Fprintf(w, "data: %d\n\n", count)
		w.(http.Flusher).Flush()
	}
}

func registerUser(user string) {
	if _, ok := users[user]; !ok {
		// user is the ip address
		users[user] = true
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)
	registerUser(ip)

	w.Write([]byte(`<!DOCTYPE html>
<html><body>
<h1 id="count">0</h1>
<script>
const es = new EventSource('/stream');
es.onmessage = e => document.getElementById('count').innerText = e.data;
</script></body></html>`))
}

// getIP extracts the real IP address from the request
func getIP(r *http.Request) string {
	// Check X-Forwarded-For header (if behind proxy/load balancer)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		// X-Forwarded-For can be comma-separated, take first IP
		return forwarded
	}

	// Check X-Real-IP header (another common proxy header)
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr (format: "IP:port")
	// Strip port if you only want the IP
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
