# Python Flask 애플리케이션 예제

FROM python:3.9-slim

# 메타데이터 추가
LABEL maintainer="instructor@example.com"
LABEL version="1.0"
LABEL description="Python Flask 애플리케이션 예제"

# 환경 변수 설정
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV APP_HOME=/app

# 작업 디렉토리 설정
WORKDIR ${APP_HOME}

# 필요한 패키지 설치
RUN pip install --no-cache-dir flask

# 애플리케이션 파일 복사
COPY app.py .
COPY requirements.txt .
COPY templates/ templates/

# 의존성 설치
RUN pip install --no-cache-dir -r requirements.txt

# 포트 노출
EXPOSE 5000

# 실행 명령어
CMD ["python", "app.py"]
