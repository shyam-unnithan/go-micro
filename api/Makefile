build:
	mkdir bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/api ./cmd/main.go
	cp -rf ./config ./bin
run:
	docker run -e MICRO_REGISTRY=mdns api

image:	build
	docker build -t api .

clean:
	-rm -rf ./bin
	-docker rmi api:latest
