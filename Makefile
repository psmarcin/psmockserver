NAME=psmockserver
IMAGE_TAG=gcr.io/$(NAME)

dev:
	realize start --name=$(NAME)

dependencies:
	go mod download

test: dependencies
	go test ./pkg/...

build: dependencies
	go build main.go

docker-build:
	docker build -t $(IMAGE_TAG) .

docker-run:
	docker run -p 8080:8080 $(IMAGE_TAG)

debug:
	dlv debug --headless --listen=:2345 --log --api-version 2

release: mock-test
	goreleaser --rm-dist

# Mock tests
mock-test:
	./test/setup.sh
