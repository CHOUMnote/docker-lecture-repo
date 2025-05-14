# Chapter 2: Docker 기본 명령어

## 학습 목표

- Docker 이미지와 컨테이너의 기본 개념 이해
- 기본적인 Docker 명령어 사용법 학습
- Docker Hub에서 이미지 검색 및 사용 방법 익히기
- 컨테이너 생명주기 관리하기

## 2.1 Docker 이미지 기본 명령어

### 이미지 검색

```bash
docker search nginx
```

### 이미지 다운로드

```bash
docker pull nginx:latest
```

### 다운로드된 이미지 목록 확인

```bash
docker images
# 또는
docker image ls
```

### 이미지 상세 정보 확인

```bash
docker image inspect nginx
```

### 이미지 삭제

```bash
docker rmi nginx
# 또는
docker image rm nginx
```

## 2.2 Docker 컨테이너 기본 명령어

### 컨테이너 생성 및 시작

```bash
# 포그라운드로 실행
docker run -p 80:80 nginx

# 백그라운드로 실행
docker run -d -p 80:80 nginx
```

주요 옵션:

- `-d`: 백그라운드 모드 (detached mode)
- `-p`: 포트 매핑 (host:container)
- `-v`: 볼륨 마운트
- `-e`: 환경 변수 설정
- `--name`: 컨테이너 이름 지정
- `--rm`: 컨테이너 종료 시 자동 삭제

### 실행 중인 컨테이너 목록 확인

```bash
docker ps
# 모든 컨테이너 확인 (중지된 컨테이너 포함)
docker ps -a
```

### 컨테이너 로그 확인

```bash
docker logs [컨테이너 ID 또는 이름]
# 실시간 로그 확인
docker logs -f [컨테이너 ID 또는 이름]
```

### 컨테이너 중지, 시작, 재시작

```bash
docker stop [컨테이너 ID 또는 이름]
docker start [컨테이너 ID 또는 이름]
docker restart [컨테이너 ID 또는 이름]
```

### 컨테이너 삭제

```bash
docker rm [컨테이너 ID 또는 이름]
# 실행 중인 컨테이너 강제 삭제
docker rm -f [컨테이너 ID 또는 이름]
```

### 컨테이너 내부 명령 실행

```bash
docker exec -it [컨테이너 ID 또는 이름] bash
```

## 2.3 실습 예제

### 예제 1: Nginx 웹 서버 실행하기

```bash
# Nginx 이미지 다운로드
docker pull nginx:latest

# 컨테이너 실행 (포트 매핑: 호스트의 8080 -> 컨테이너의 80)
docker run -d --name my-nginx -p 8080:80 nginx

# 브라우저에서 http://localhost:8080 접속하여 확인
# 컨테이너 로그 확인
docker logs my-nginx
```

### 예제 2: 환경 변수를 사용한 MySQL 컨테이너 실행

```bash
docker run -d \
  --name my-mysql \
  -e MYSQL_ROOT_PASSWORD=my-secret-pw \
  -e MYSQL_DATABASE=mydb \
  -p 3306:3306 \
  mysql:8
```

## 2.4 유용한 Docker 명령어

### 시스템 정보 확인

```bash
# Docker 시스템 정보 확인
docker info

# 디스크 사용량 확인
docker system df

# 미사용 리소스(이미지, 컨테이너, 볼륨) 정리
docker system prune
```

### 컨테이너 리소스 사용량 확인

```bash
docker stats
```

## 다음 단계

[Chapter 3: Dockerfile 작성하기](../chapter3-dockerfile/README.md)로 넘어가서 직접 Docker 이미지를 빌드하는 방법을 학습해 보세요.
