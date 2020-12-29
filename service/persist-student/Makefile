build:
# Ensure you have used export GOPATH
	echo "GOPATH: ${GOPATH}"
	mkdir -p ${GOPATH}/bin/service/persist-student/
	go build -o ${GOPATH}/bin/service/persist-student/cmd ./cmd/main.go
	cp -rf ./config ${GOPATH}/bin/service/persist-student

compile:
# Ensure you have used export GOPATH
	$(MAKE) clean
	mkdir -p bin/service/persist-student
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -installsuffix cgo -o bin/service/persist-student/cmd ./cmd/main.go
	cp -rf ./config bin/service/persist-student
	cp -rf bin/service/persist-student/cmd ${GOPATH}/bin/service/persist-student

run:
	$(MAKE) image
	docker run --name persist-student --network gomicro -d persist-student

image:
	$(MAKE) compile	
	docker build -t persist-student .

clean:
	rm -rf ${GOPATH}/src/service/persist-student/bin/persist-student 
	rm -rf ./bin
	
dockerclean:
	docker rm persist-student

