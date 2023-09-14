FROM golang:1.21-alpine

ENV GIN_MODE=debug
ENV PORT=5000

RUN apk update && apk add --no-cache git bash openssh

WORKDIR /app

COPY go.mod go.sum ./
        
RUN go mod download

COPY . .
        
# The issue is, we are connecting the /app folder to local folder and when the executable is created
# ...it's done for macOS instead of ubuntu. We need a folder that's not mapped to local storage.

RUN go build -o /go/main .

ENTRYPOINT ["/go/main"]