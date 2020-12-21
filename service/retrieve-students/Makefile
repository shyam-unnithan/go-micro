build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o retrieve-student-service ./cmd/main.go

run:
	docker run -e MICRO_REGISTRY=mdns retrieve-student-service

image:	build
	docker build -t retrieve-student-service .

clean:
	-rm -rf ./retrieve-student-service
	-docker rmi retrieve-student-service:latest
