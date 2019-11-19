.PHONY: clean build

build:
	cd ./hello/ && GOOS=linux GOARCH=amd64 go build

clean:
	rm -rf ./hello/hello

