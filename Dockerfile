FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install -v github.com/fdaines/arch-go@latest
RUN go install github.com/fdaines/spm-go@latest
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy