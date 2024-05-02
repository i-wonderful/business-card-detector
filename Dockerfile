FROM golang:1.20.1

RUN mkdir /app
WORKDIR /app

# Копируем все файлы проекта в рабочую директорию
#COPY . /app
COPY ./cmd /app/cmd
COPY ./config /app/config
COPY ./internal /app/internal
COPY ./template /app/template
COPY ./lib /app/lib
COPY ./onnx /app/onnx
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
RUN mkdir /app/tmp
RUN mkdir /app/storage

# Устанавливаем пакеты, необходимые для использования gosseract
RUN apt-get update && \
    apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr tesseract-ocr-eng tesseract-ocr-rus

ENV GO111MODULE=on
RUN go build -o main ./cmd

EXPOSE 8080
CMD ["/app/main"]

