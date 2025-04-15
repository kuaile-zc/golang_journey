DATETIME=`date +'%y%m%d%H%M%S'`
BUILD_DIR=./bin

build:
	rm -rf ${BUILD_DIR}/get_url
	GO111MODULE=on go build -mod=vendor -o ${BUILD_DIR}/get_url/get_url main/get_url/main.go
	rm -rf ${BUILD_DIR}/web_server
	GO111MODULE=on go build -mod=vendor -o ${BUILD_DIR}/web_server/web_server main/web_server/main.go

compile:
	go mod tidy && go mod vendor

clean:
	rm -rf ${BUILD_DIR}



