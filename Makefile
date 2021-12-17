#  "make server  -> runs server on port 8081"
#  "make client  -> runs client and connects to 127.0.0.1:8081"
#  "make install -> installs the program"
#  "make build   -> compiles the code and build a bin file"
#  "make compile -> compiles the code for all OS and build bin file for them"

server:
	go run main.go server 8081

client:
	go run main.go client 127.0.0.1:8081

install:
	go install go-cloud-calculator

build:
	go build -o bin/go-cloud-calculator main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/go-cloud-calculator-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/go-cloud-calculator-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/go-cloud-calculator-freebsd-386 main.go
