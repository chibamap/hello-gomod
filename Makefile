.PHONY: clean build

build:
	go fmt && GOOS=linux GOARCH=amd64 go build -o build/hello-gomod

clean:
	rm -rf ./build/hello-gomod

