build:
# Ensure you have used export GOPATH
	echo "GOPATH: ${GOPATH}"
	mkdir -p ${GOPATH}/bin/service/event/
	go build -o ${GOPATH}/bin/service/event/cmd ./cmd/main.go
	cp -rf ./config ${GOPATH}/bin/service/event

compile:
# Ensure you have used export GOPATH
	$(MAKE) clean
	mkdir -p bin/service/event
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -installsuffix cgo -o bin/service/event/cmd ./cmd/main.go
	cp -rf ./config bin/service/event
	cp -rf bin/service/event/cmd ${GOPATH}/bin/service/event

run:
#	docker run -e MICRO_REGISTRY=mdns event
	$(MAKE) image
	docker run --name event --network gomicro -d event

image:
	$(MAKE) compile	
	docker build -t event .

clean:
	rm -rf ${GOPATH}/src/service/event/bin/event 
	rm -rf ./bin
	
dockerclean:
	docker rmi event:latest

