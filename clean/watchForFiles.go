package clean

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"time"

	"github.com/radovskyb/watcher"
)

func getDownloadDir() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting the current user: %v", err)
	}

	downloadDir := filepath.Join(u.HomeDir, "Downloads")
	return downloadDir
}

func WatchForDownloads() {
	downloadDir := getDownloadDir()

	w := watcher.New()

	if err := w.AddRecursive(downloadDir); err != nil {
		log.Fatalf("Error when adding a directory to be monitored: %v", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-w.Event:
				if !ok {
					return
				}
				if event.Op == watcher.Create {
					fmt.Println("New file:", event.Path)
					checkExtension(event.Path)
				}
			case err, ok := <-w.Error:
				if !ok {
					return
				}
				log.Println("Monitoring error:", err)
			}
		}
	}()

	if err := w.Start(250 * time.Millisecond); err != nil {
		log.Fatalf("Error starting the monitoring: %v", err)
	}

	select {}
}
