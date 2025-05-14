# 간단한 Nginx 이미지 예제

FROM nginx:alpine

# 메타데이터 추가
LABEL maintainer="instructor@example.com"
LABEL version="1.0"
LABEL description="간단한 Nginx 커스텀 이미지"

# 작업 디렉토리 설정
WORKDIR /usr/share/nginx/html

# 기본 index.html 파일 대체
COPY index.html .

# 포트 노출
EXPOSE 80

# 기본 명령어 (Nginx 이미지에 이미 정의되어 있으므로 필요 없음)
# CMD ["nginx", "-g", "daemon off;"]
