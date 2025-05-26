package main

import (
	"log"
	"net/http"
)

func main() {
	// 현재 디렉토리에서 파일 제공
	fs := http.FileServer(http.Dir("."))

	// 모든 경로에 파일 서버 핸들러 연결
	http.Handle("/", fs)

	log.Println("🚀 Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
