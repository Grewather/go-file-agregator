package clean

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"time"

	"github.com/Grewather/go-file-agregator/system"
	"github.com/radovskyb/watcher"
)

func getDownloadDir() string {
	u, err := user.Current()
	if err != nil {
		system.ShowMessageBox("Error getting the current user: " + err.Error())
	}

	downloadDir := filepath.Join(u.HomeDir, "Downloads")
	return downloadDir
}

func WatchForDownloads() {
	downloadDir := getDownloadDir()

	w := watcher.New()

	if err := w.AddRecursive(downloadDir); err != nil {
		system.ShowMessageBox("Error when adding a directory to be monitored" + err.Error())
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
		system.ShowMessageBox("Error starting the monitoring"+  err.Error())
	}

	select {}
}
