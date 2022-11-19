# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /honest-truth-api

EXPOSE 80
CMD [ "/honest-truth-api"]