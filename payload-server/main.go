package main

import (
        "net/http"
)

var (
        p1KB   = fillData(1 * 1024)
        p10KB  = fillData(10 * 1024)
        p100KB = fillData(100 * 1024)
        p1MB   = fillData(1 * 1024 * 1024)
        p10MB  = fillData(10 * 1024 * 1024)
)

func fillData(size int) []byte {
        b := make([]byte, size)
        for i := range b {
                b[i] = 'X'
        }
        return b
}

func handler(payload []byte) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/octet-stream")
                //w.Header().Set("Content-Length", contentLength)
                //w.WriteHeader(http.StatusOK)
                w.Write(payload)
        }
}

func main() {
        mux := http.NewServeMux()
        mux.HandleFunc("/1kb", handler(p1KB))
        mux.HandleFunc("/10kb", handler(p10KB))
        mux.HandleFunc("/100kb", handler(p100KB))
        mux.HandleFunc("/1mb", handler(p1MB))
        mux.HandleFunc("/10mb", handler(p10MB))

        port := "8080"
        http.ListenAndServe(":"+port, mux)
}
