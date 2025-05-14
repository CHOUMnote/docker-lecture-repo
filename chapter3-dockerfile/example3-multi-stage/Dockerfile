# 1단계: 빌드 단계
FROM golang:1.20 AS builder

# 작업 디렉토리 설정
WORKDIR /app

# go.mod와 go.sum 파일 생성
RUN go mod init example3-multi-stage && \
    go mod tidy

# 소스 코드 복사
COPY main.go .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 go build -o main .

# 2단계: 실행 단계
FROM debian:bookworm-slim

# 필요한 파일 복사
COPY --from=builder /app/main /app/main

# 실행 디렉토리 설정
WORKDIR /app

# 포트 노출
EXPOSE 8080

# 애플리케이션 실행
CMD ["./main"]