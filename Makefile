.PHONY: clean build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/hello-gomod

clean:
	rm -rf ./hello/hello

