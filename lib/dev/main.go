package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	root, err := findModuleRoot()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	appPort, err := freePort()
	if err != nil {
		log.Fatal(err)
	}

	buildDir, err := os.MkdirTemp("", "constellation-dev-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(buildDir)

	serverRunner := newRunner(root, buildDir, appPort)
	binary, err := serverRunner.build()
	if err != nil {
		log.Fatal(err)
	}
	if err := serverRunner.start(binary); err != nil {
		log.Fatal(err)
	}
	defer serverRunner.stop()

	if err := waitForPort(appPort, 10*time.Second); err != nil {
		log.Fatal(err)
	}

	hub := newReloadHub()
	server := &http.Server{
		Addr:              net.JoinHostPort("", port),
		Handler:           newProxy(appPort, hub),
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("hot reload ready at http://localhost:%s (watching .go, .css, .md)", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	watcher, err := newWatcher(root)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if watcher.changed() {
			log.Println("change detected, rebuilding...")

			binary, err := serverRunner.build()
			if err != nil {
				log.Printf("%v (keeping the previous server)", err)
				continue
			}
			if err := serverRunner.restart(binary); err != nil {
				log.Printf("restart failed: %v", err)
				continue
			}

			hub.broadcast()
			log.Println("reloaded")
		}

		select {
		case <-ctx.Done():
			shutdown, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			_ = server.Shutdown(shutdown)
			return
		default:
		}
	}
}

func findModuleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found in current directory or parents")
		}
		dir = parent
	}
}

func freePort() (string, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer listener.Close()

	_, port, err := net.SplitHostPort(listener.Addr().String())
	return port, err
}
