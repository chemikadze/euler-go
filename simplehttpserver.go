package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"log"
)

func RespondCode(w http.ResponseWriter, code int, body string) {
	w.WriteHeader(code)
	w.Write([]byte(body))
}

func formatDirectoryPlain(files []os.FileInfo) (contentType string, responseBody string) {
	response := ""
	for _, file := range files {
		response += file.Name() + "\n"
	}
	return "text/plain", response
}

func formatDirectoryHtml(files []os.FileInfo) (contentType string, responseBody string) {
	items := ""
	for _, file := range files {
		items += "<li>" + file.Name() + "</li>"
	}
	contentType = "text/html"
	responseBody = "<html><head><title>SimpleHttpServer-go</title></head><body><ul>" + items + "</ul></body></html>"
	return
}

type dirFormat func([]os.FileInfo) (string, string)

var acceptedContentTypes = map[string]dirFormat{
	"text/plain": formatDirectoryPlain,
	"text/html":  formatDirectoryHtml,
}

// yes, i know 'bout http.FileServer
func FileServer(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	path := req.URL.Path
	dir, _ := os.Getwd()
	absPath, _ := filepath.Abs(filepath.Join(dir, path))
	stat, err := os.Stat(absPath)
	if err != nil {
		RespondCode(w, http.StatusInternalServerError, err.Error())
		return
	}
	if stat.IsDir() {
		files, err := ioutil.ReadDir(absPath)
		if err != nil {
			RespondCode(w, http.StatusInternalServerError, err.Error())
			return
		}
		accept := req.Header.Get("Accept")
		queryOverride := req.URL.Query().Get("accept")
		if queryOverride != "" {
			accept = queryOverride
		}
		formatter := acceptedContentTypes[accept]
		if formatter == nil {
			formatter = formatDirectoryPlain
		}
		contentType, responseBody := formatter(files)
		w.Header().Add("Content-Type", contentType)
		w.Write([]byte(responseBody))
	} else {
		data, err := ioutil.ReadFile(absPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte(data))
	}
}

func main() {
	port := 8080
	if len(os.Args) > 1 {
		parsedPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err.Error())
		}
		port = parsedPort
	}
	http.HandleFunc("/", FileServer)
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), nil))
}
