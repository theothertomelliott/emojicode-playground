.PHONY: clean build docker

clean:
	rm -rf build

build:
	GOOS=linux GOARCH=amd64 go build -o build/emojicode-playground ./cmd/emojicode-playground

docker: build
	docker build -t theothertomelliott/emojicode-playground:latest ./build -f Dockerfile

compose: build
	docker-compose up --build