.PHONY: all test clean

install:
	go get .
	
build-win:
	go build -o ./bin/converter.exe ./main.go

build:
	go build -o ./bin/converter ./main.go && chmod +x ./bin/converter