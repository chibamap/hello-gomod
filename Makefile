.PHONY: clean build

build:
	go fmt && GOOS=linux GOARCH=amd64 go build -o bin/hello-gomod

clean:
	rm -rf ./bin/hello-gomod

