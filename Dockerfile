FROM golang:1.20-alpine as base

FROM base as dev

WORKDIR /opt/app
COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]
