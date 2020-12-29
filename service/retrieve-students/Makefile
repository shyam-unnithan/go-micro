build:
# Ensure you have used export GOPATH
	echo "GOPATH: ${GOPATH}"
	mkdir -p ${GOPATH}/bin/service/retrieve-students/
	go build -o ${GOPATH}/bin/service/retrieve-students/cmd ./cmd/main.go
	cp -rf ./config ${GOPATH}/bin/service/retrieve-students

compile:
# Ensure you have used export GOPATH
	$(MAKE) clean
	mkdir -p bin/service/retrieve-students
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -installsuffix cgo -o bin/service/retrieve-students/cmd ./cmd/main.go
	cp -rf ./config bin/service/retrieve-students
	cp -rf bin/service/retrieve-students/cmd ${GOPATH}/bin/service/retrieve-students

run:
#	docker run -e MICRO_REGISTRY=mdns retrieve-students
	$(MAKE) image
	docker run --name retrieve-students --network gomicro -d retrieve-students

image:
	$(MAKE) compile
	docker build -t retrieve-students .

clean:
	rm -rf ${GOPATH}/src/service/retrieve-students/bin/retrieve-students 
	rm -rf ./bin

