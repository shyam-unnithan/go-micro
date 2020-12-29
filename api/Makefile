build:
# Ensure you have used export GOPATH
	echo "GOPATH: ${GOPATH}"
	mkdir ${GOPATH}/bin/api
	go build -o ${GOPATH}/bin/api/cmd ./cmd/main.go
	cp -rf ./config ${GOPATH}/bin/api

compile:
# Ensure you have used export GOPATH
	$(MAKE) clean
	mkdir -p bin/api
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -installsuffix cgo -o bin/api/cmd ./cmd/main.go
	cp -rf ./config bin/api
	cp -rf cmd ${GOPATH}/bin/api

run:
#	docker run -e MICRO_REGISTRY=mdns api
	$(MAKE) image
	$(MAKE) dockerclean
	docker run --name api --network gomicro api

image:	
	$(MAKE) compile
	docker build -t api .

clean:
	rm -rf ${GOPATH}/bin/api 
	rm -rf ./bin
	
dockerclean:
	docker rm api

