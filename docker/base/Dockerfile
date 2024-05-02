FROM golang:latest

# Устанавливаем пакеты, необходимые для использования gosseract
RUN apt-get update && \
    apt-get -y install libleptonica-dev && \
    apt-get -y install tesseract-ocr && \
    apt-get -y install libtesseract-dev
RUN apt-get install -y \
  tesseract-ocr-eng \
  tesseract-ocr-rus

CMD ["/app/main"]
