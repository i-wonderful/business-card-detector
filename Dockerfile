# 1. Build backend
FROM golang:1.20-alpine AS go-build

RUN mkdir /app
WORKDIR /app

COPY ./cmd /app/cmd
COPY ./config /app/config
COPY ./internal /app/internal
COPY ./template /app/template
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

ENV GO111MODULE=on
RUN go build -o main ./cmd

# 2. Run paddle ocr ---------------------------
FROM paddlepaddle/paddle:2.6.1
#FROM paddlepaddle/paddle:2.6.1-jupyter

RUN mkdir -p /root/.paddleocr/whl
COPY ./lib/paddleocr/whl /root/.paddleocr/whl
#RUN apt-get update && apt-get install libgomp1 libgl1-mesa-glx libgtk2.0-0

RUN pip install "paddleocr>=2.0.1"

#RUN python --version
#RUN pip show paddleocr

RUN mkdir -p /app/storage
WORKDIR /app

# Copy the Golang binary
COPY --from=go-build /app/main /app
COPY --from=go-build /app/config /app/config
COPY --from=go-build /app/internal/service/text_recognize/paddleocr/run.py /app/internal/service/text_recognize/paddleocr/run.py
COPY  --from=go-build /app/template /app/template

EXPOSE 8080
CMD ["/app/main"]
