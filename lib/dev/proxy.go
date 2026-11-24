package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	reloadPath   = "/__constellation_reload"
	reloadScript = `<script>(function(){var s=new EventSource("` + reloadPath + `");s.addEventListener("reload",function(){location.reload()})})()</script>`
)

type reloadHub struct {
	mu      sync.Mutex
	clients map[chan struct{}]struct{}
}

func newReloadHub() *reloadHub {
	return &reloadHub{clients: make(map[chan struct{}]struct{})}
}

func (h *reloadHub) broadcast() {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients {
		select {
		case client <- struct{}{}:
		default:
		}
	}
}

func (h *reloadHub) serveHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	client := make(chan struct{}, 1)
	h.mu.Lock()
	h.clients[client] = struct{}{}
	h.mu.Unlock()
	defer func() {
		h.mu.Lock()
		delete(h.clients, client)
		h.mu.Unlock()
	}()

	fmt.Fprint(w, ": connected\n\n")
	flusher.Flush()

	keepAlive := time.NewTicker(15 * time.Second)
	defer keepAlive.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-client:
			fmt.Fprint(w, "event: reload\ndata: {}\n\n")
			flusher.Flush()
		case <-keepAlive.C:
			fmt.Fprint(w, ": ping\n\n")
			flusher.Flush()
		}
	}
}

func newProxy(port string, hub *reloadHub) http.Handler {
	target := &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("127.0.0.1", port),
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = injectReloadScript

	mux := http.NewServeMux()
	mux.HandleFunc(reloadPath, hub.serveHTTP)
	mux.Handle("/", proxy)
	return mux
}

func injectReloadScript(response *http.Response) error {
	if !strings.Contains(response.Header.Get("Content-Type"), "text/html") {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	_ = response.Body.Close()

	var patched []byte
	if index := bytes.LastIndex(body, []byte("</body>")); index >= 0 {
		patched = make([]byte, 0, len(body)+len(reloadScript))
		patched = append(patched, body[:index]...)
		patched = append(patched, reloadScript...)
		patched = append(patched, body[index:]...)
	} else {
		patched = append(body, reloadScript...)
	}

	response.Body = io.NopCloser(bytes.NewReader(patched))
	response.ContentLength = int64(len(patched))
	response.Header.Set("Content-Length", strconv.Itoa(len(patched)))
	return nil
}
