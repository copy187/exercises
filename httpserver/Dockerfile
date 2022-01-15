FROM golang:1.17-alpine as build

ENV GOPROXY=https://goproxy.cn

RUN apk add --nocache git
COPY . /httpserver
WORKDIR /httpserver
RUN go mod tidy && mkdir -p bin/amd64 && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

FROM scratch

COPY --from=build /httpserver/bin/amd64/httpserver /httpserver
ENTRYPOINT ["/httpserver"]





