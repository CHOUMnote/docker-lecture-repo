# Chapter 3: Dockerfile 작성하기

## 학습 목표
- Dockerfile의 기본 구조와 문법 이해
- 다양한 Dockerfile 지시어(instruction) 사용법 학습
- 효율적인 Docker 이미지 빌드 방법 이해
- 멀티 스테이지 빌드 활용하기

## 3.1 Dockerfile 기본 개념

Dockerfile은 Docker 이미지를 빌드하기 위한 스크립트입니다. 이 파일에는 이미지를 구성하는 일련의 명령어가 포함되어 있습니다.

기본 구조:
```dockerfile
# 베이스 이미지 지정
FROM base_image:tag

# 메타데이터 설정
LABEL maintainer="your-email@example.com"
LABEL version="1.0"

# 환경 변수 설정
ENV APP_HOME=/app

# 작업 디렉토리 설정
WORKDIR /app

# 파일 복사
COPY . .

# 명령어 실행
RUN apt-get update && apt-get install -y package-name

# 포트 노출
EXPOSE 8080

# 컨테이너 실행 시 실행할 명령어
CMD ["executable", "param1", "param2"]
```

## 3.2 주요 Dockerfile 지시어

### FROM
베이스 이미지를 지정합니다. 모든 Dockerfile은 FROM 명령어로 시작해야 합니다.
```dockerfile
FROM ubuntu:20.04
```

### LABEL
이미지에 메타데이터를 추가합니다.
```dockerfile
LABEL maintainer="your-email@example.com"
LABEL version="1.0"
LABEL description="This is my custom image"
```

### ENV
환경 변수를 설정합니다.
```dockerfile
ENV APP_HOME=/app
ENV NODE_ENV=production
```

### WORKDIR
작업 디렉토리를 설정합니다. 이후의 명령어(RUN, CMD, ENTRYPOINT, COPY, ADD)는 이 디렉토리를 기준으로 실행됩니다.
```dockerfile
WORKDIR /app
```

### COPY와 ADD
파일이나 디렉토리를 이미지에 복사합니다. ADD는 추가 기능(URL에서 다운로드, tar 파일 자동 압축 해제)을 제공하지만, 일반적으로 단순한 복사는 COPY를 사용하는 것이 좋습니다.
```dockerfile
COPY . .
COPY package.json .
ADD http://example.com/file.tar.gz /tmp/
```

### RUN
이미지 빌드 과정에서 명령어를 실행합니다.
```dockerfile
RUN apt-get update && apt-get install -y nodejs npm
```

### EXPOSE
컨테이너가 런타임에 리스닝할 포트를 지정합니다. (실제 포트 연결은 docker run -p로 해야 함)
```dockerfile
EXPOSE 8080
```

### CMD와 ENTRYPOINT
컨테이너가 시작될 때 실행할 명령어를 지정합니다.
- CMD: 실행할 기본 명령어 (docker run 시 override 가능)
- ENTRYPOINT: 컨테이너가 실행될 때 항상 실행되는 명령어

```dockerfile
# CMD 방식 1: 셸 형식
CMD npm start

# CMD 방식 2: 실행 형식 (권장)
CMD ["npm", "start"]

# ENTRYPOINT
ENTRYPOINT ["node", "app.js"]
```

## 3.3 Dockerfile 예제

이 챕터의 세 가지 예제를 확인해 보세요:

1. [예제 1: 간단한 Dockerfile](./example1-simple/)
2. [예제 2: Python 애플리케이션](./example2-python-app/)
3. [예제 3: 멀티 스테이지 빌드](./example3-multi-stage/)

## 3.4 이미지 빌드 및 관리

### 이미지 빌드
```bash
docker build -t myapp:1.0 .
```

주요 옵션:
- `-t` 또는 `--tag`: 이미지 이름과 태그 지정
- `-f` 또는 `--file`: Dockerfile 위치 지정 (기본값: ./Dockerfile)
- `--no-cache`: 캐시 사용하지 않고 빌드

### 이미지 히스토리 확인
```bash
docker history myapp:1.0
```

## 3.5 Docker 빌드 컨텍스트

Docker 빌드 명령을 실행할 때, 현재 디렉토리(또는 지정된 디렉토리)의 모든 파일과 디렉토리가 빌드 컨텍스트로 Docker 데몬에 전송됩니다. 불필요한 파일이 포함되지 않도록 .dockerignore 파일을 사용하는 것이 좋습니다.

예시 .dockerignore:
```
node_modules
npm-debug.log
.git
.gitignore
```

## 3.6 빌드 모범 사례

1. 가능한 작은 베이스 이미지 사용하기
2. 멀티 스테이지 빌드 활용하기
3. 레이어 수 최소화하기 (RUN 명령어 체이닝)
4. .dockerignore 파일 사용하기
5. 적절한 태그 사용하기
6. 불필요한 패키지 설치 피하기

## 다음 단계
[Chapter 4: Docker 볼륨과 네트워크](../chapter4-volumes-networks/README.md)로 넘어가서 Docker의 데이터 관리와 네트워킹에 대해 학습해 보세요.
