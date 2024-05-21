FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y tesseract-ocr \
        tesseract-ocr-eng \
        tesseract-ocr-rus && \
    apt-get clean

CMD ["tesseract", "--version"]