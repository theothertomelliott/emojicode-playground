.PHONY: clean build docker

REPOSITORY=theothertomelliott
VERSION=0.1

clean:
	rm -rf build

build:
	GOOS=linux GOARCH=amd64 go build -o build/emojicode-playground ./cmd/emojicode-playground
	cp -r static ./build/static

docker: build
	docker build -t emojicode-playground:latest ./build -f Dockerfile

push: docker
	docker tag emojicode-playground:latest ${REPOSITORY}/emojicode-playground:${VERSION} 
	docker push ${REPOSITORY}/emojicode-playground:${VERSION}

compose: build
	docker-compose up --build