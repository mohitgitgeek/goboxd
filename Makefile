IMAGE ?= goboxd:dev
CONTAINER_NAME ?= goboxd
ARCHIVE ?= goboxd-image.tar

.PHONY: build run test integration load lint save fmt

build:
	docker build -t $(IMAGE) -f Dockerfile .

run:
	docker compose up --build $(CONTAINER_NAME)

test: build
	docker run --rm --entrypoint /bin/bash $(IMAGE) -c "go test ./..."

integration: build
	docker run --rm --entrypoint /bin/bash $(IMAGE) -c "go test -tags=integration ./integration/... && ./tests/e2e/run_python.sh"

lint: build
	docker run --rm --entrypoint /bin/bash $(IMAGE) -c "golangci-lint run ./... && staticcheck ./..."

fmt: build
	docker run --rm --entrypoint /bin/bash $(IMAGE) -c "go fmt ./..."

save: build
	docker save -o $(ARCHIVE) $(IMAGE)

load:
	docker load -i $(ARCHIVE)
