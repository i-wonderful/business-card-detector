# Базовый образ для Golang
FROM cardbase:latest

# Создаем и переключаемся в рабочую директорию
RUN mkdir /app
WORKDIR /app

# Копируем все файлы проекта в рабочую директорию
#COPY . /app
COPY ../../cmd /app/cmd
COPY ../../config /app/config
COPY ../../internal /app/internal
COPY ../../lib /app/lib
COPY ../../onnx /app/onnx
COPY ../../go.mod /app/go.mod
COPY ../../go.sum /app/go.sum
RUN mkdir /app/tmp

# Устанавливаем пакеты, необходимые для использования gosseract
#RUN apt-get update && \
#    apt-get -y install libleptonica-dev && \
#    apt-get -y install tesseract-ocr && \
#    apt-get -y install libtesseract-dev
#RUN apt-get install -y \
#  tesseract-ocr-eng \
#  tesseract-ocr-rus


ENV GO111MODULE=on
RUN go build -o main ./cmd

EXPOSE 8080
CMD ["/app/main"]
