# Chapter 1: Docker 소개

## 학습 목표

- Docker의 개념과 작동 원리 이해
- 컨테이너와 가상 머신의 차이점 이해
- Docker의 주요 구성 요소 파악
- Docker 설치 및 설정

## 1.1 Docker란 무엇인가?

Docker는 애플리케이션을 개발, 배포, 실행하기 위한 오픈소스 플랫폼입니다. Docker를 사용하면 애플리케이션을 인프라로부터 분리하여 소프트웨어를 빠르게 제공할 수 있습니다.

Docker의 주요 이점:

- 일관된 환경 제공
- 빠른 배포 및 확장
- 효율적인 리소스 사용
- 개발과 운영 환경의 간극 해소

## 1.2 컨테이너 vs 가상 머신

### 가상 머신 (VM)

- 하이퍼바이저를 통해 여러 OS를 실행
- 각 VM은 자체 OS 커널을 포함
- 리소스 사용량이 많음
- 시작 시간이 느림

### 컨테이너

- 호스트 OS 커널을 공유
- 필요한 바이너리와 라이브러리만 포함
- 경량화되어 리소스 효율적
- 빠르게 시작 및 중지

## 1.3 Docker 아키텍처

Docker는 클라이언트-서버 아키텍처를 사용합니다:

1. **Docker 클라이언트**: 사용자가 Docker 명령어를 입력하는 인터페이스
2. **Docker 데몬(dockerd)**: Docker API 요청을 수신하고 컨테이너, 이미지, 네트워크 등을 관리
3. **Docker 레지스트리**: Docker 이미지를 저장하고 배포하는 저장소 (Docker Hub 등)

## 1.4 Docker 설치하기

### Windows

1. [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop) 다운로드
2. 설치 프로그램 실행 및 지시 따르기
3. WSL2 설정 필요 (Windows 10 이상)

### macOS

1. [Docker Desktop for Mac](https://www.docker.com/products/docker-desktop) 다운로드
2. 설치 프로그램 실행 및 지시 따르기

### Linux (Ubuntu)

```bash
# 필요한 패키지 설치
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common

# Docker의 공식 GPG 키 추가
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Docker 저장소 추가
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Docker 설치
sudo apt-get update
sudo apt-get install docker-ce

# Docker 그룹에 사용자 추가 (sudo 없이 사용하기 위함)
sudo usermod -aG docker $USER
```

## 1.5 Docker 버전 확인

설치가 완료되면 다음 명령어로 Docker 버전을 확인할 수 있습니다:

```bash
docker version
docker info
```

## 다음 단계

[Chapter 2: Docker 기본 명령어](../chapter2-basic-commands/README.md)로 넘어가서 Docker의 기본 명령어들을 학습해 보세요.
