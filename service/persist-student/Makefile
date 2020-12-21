build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o persist-student-service ./cmd/main.go

run:
	docker run -e MICRO_REGISTRY=mdns persist-student-service

image:	build
	docker build -t persist-student-service .

clean:
	-rm -rf ./persist-student-service
	-docker rmi persist-student-service:latest
