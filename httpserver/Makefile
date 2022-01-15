export tag=v1.0

root:
	export ROOT=github.com/copy187/exercises

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release:
	echo "building httpserver container"
    docker build -t copy187/httpserver:${tag} .

push: release
	echo "pushing copy187/httpserver"
	docker push copy187/httpserver:v1.0