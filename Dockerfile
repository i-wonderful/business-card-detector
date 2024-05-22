FROM golang:latest

RUN mkdir /app
WORKDIR /app

# Копируем все файлы проекта в рабочую директорию
COPY ./cmd /app/cmd
COPY ./config /app/config
COPY ./internal /app/internal
COPY ./template /app/template
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
RUN mkdir /app/storage

RUN mkdir -p /root/.paddleocr/whl
COPY ./lib/paddleocr/whl /root/.paddleocr/whl

# Обновляем список пакетов и устанавливаем Python и pip
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv \
     libglib2.0-0 libgl1-mesa-glx

# Create a virtual environment
RUN python3 -m venv /app/venv

# Activate the virtual environment and install packages
RUN /app/venv/bin/pip install paddlepaddle paddleocr

# Set the environment variables
ENV VIRTUAL_ENV=/app/venv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

ENV GO111MODULE=on
RUN go build -o main ./cmd

# Check Python version
#RUN python3 --version

# Check versions of installed libraries
#RUN /app/venv/bin/pip show paddlepaddle paddleocr

EXPOSE 8080
CMD ["/app/main"]

