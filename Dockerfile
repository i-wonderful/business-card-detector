# 1. Build backend
FROM golang:1.20-buster AS go-build

# Установка зависимостей для CGO
RUN apt-get update && apt-get install -y gcc

RUN mkdir /app
WORKDIR /app

COPY ./cmd /app/cmd
COPY ./config /app/config
COPY ./internal /app/internal
COPY ./template /app/template
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN GOOS=linux go build -o main ./cmd

# 2. Run paddle ocr ---------------------------
FROM python:3.9-slim-buster

# Install ONNX Runtime and RapidOCR ONNX Runtime
RUN pip install onnxruntime rapidocr_onnxruntime

# Install dependencies and clean up apt lists to reduce image size
RUN apt-get update && apt-get install -y \
    libgirepository1.0-dev \
    gir1.2-gtk-3.0 \
    libcairo2-dev \
    pkg-config \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libavutil-dev \
    libxvidcore-dev \
    libx264-dev \
    libgtk-3-dev \
    libgstreamer-plugins-base1.0-dev \
    libc6 \
    gcc && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir /app
WORKDIR /app

# Copy the Golang binary
COPY --from=go-build /app/main /app
COPY --from=go-build /app/config /app/config
COPY --from=go-build /app/template /app/template
COPY ./lib /app/lib
COPY ./python /app/python
COPY ./storage /app/storage

EXPOSE 8080
CMD ["/app/main"]
