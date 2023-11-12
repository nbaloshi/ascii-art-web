FROM golang:1.20-alpine

WORKDIR /app


COPY . .

RUN go build -o app

LABEL unwanted=/temp
LABEL intitute="Reboot01"
LABEL version="1.0"
LABEL description="This was easy"

EXPOSE 8080:8080

ENTRYPOINT ["./app"]
