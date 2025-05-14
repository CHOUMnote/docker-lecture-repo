package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 환경변수에서 포트 가져오기, 없으면 8080 사용
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 기본 핸들러 설정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("요청 받음: %s %s", r.Method, r.URL.Path)
		fmt.Fprintf(w, "안녕하세요! Go 웹 서버입니다!\n")
	})

	// 상태 확인용 엔드포인트
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "정상 동작 중\n")
	})

	// 서버 시작
	log.Printf("서버가 :%s 포트에서 시작됩니다...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
	}
}
