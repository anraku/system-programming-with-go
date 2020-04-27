package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "world",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	mw := io.MultiWriter(os.Stdout, gzipWriter)

	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(source)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
