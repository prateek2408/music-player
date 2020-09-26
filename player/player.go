package main

import (
	"encoding/json"
	"fmt"
	"github.com/prateek2408/music-player/utils"
	"net/http"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Mode  os.FileMode
}

const (
	filePrefix = "/music/"
	root       = "./music"
)

func main() {
	fmt.Printf("\nMusic Player version alpha")
	fmt.Printf("\nIntialising")
	http.HandleFunc("/", loadMainFrame)
	http.HandleFunc(filePrefix, File)
	http.ListenAndServe(":8080", nil)

}

func loadMainFrame(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./player.html")
}

func File(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(root, r.URL.Path[len(filePrefix):])
	stat, err := os.Stat(path)
	utils.Chk(err)
	if stat.IsDir() {
		serveDir(w, r, path)
	}
	http.ServeFile(w, r, path)
}

func serveDir()
