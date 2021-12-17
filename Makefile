#  "make server  -> runs server on port 8081"
#  "make client  -> runs client and connects to 127.0.0.1:8081"
#  "make install -> installs the program"
#  "make build   -> compiles the code and build a bin file"
#  "make compile -> compiles the code for all OS and build bin file for them"
#  "make add-lib -> adds required libs"

server:
	go run cmd/cloud-calculator/main.go server 8081

client:
	go run cmd/cloud-calculator/main.go client 127.0.0.1:8081

install:
	go install ./cmd/cloud-calculator

build:
	go build -o bin/cloud-calculator cmd/cloud-calculator/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/cloud-calculator-linux-arm cmd/cloud-calculator/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/cloud-calculator-linux-arm64 cmd/cloud-calculator/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/cloud-calculator-freebsd-386 cmd/cloud-calculator/main.go

add-lib:
	go get -u github.com/spf13/cobra/cobra
	go get -u github.com/spf13/cobra/viper
