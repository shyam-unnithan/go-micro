build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64	go build -a -installsuffix cgo -o event-service

run:
	docker run -e MICRO_REGISTRY=mdns event-service

image:	build
	docker build -t event-service .

clean:
	-rm -rf ./event-service
	-docker rmi event-service:latest
