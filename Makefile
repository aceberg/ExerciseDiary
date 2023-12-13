PKG_NAME=AppTemplate
USR_NAME=aceberg

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/$(USR_NAME)/$(PKG_NAME) && \
	go mod tidy

run:
	cd cmd/$(PKG_NAME)/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/$(PKG_NAME)/ && \
	CGO_ENABLED=0 go build -o ../../tmp/$(PKG_NAME) .

docker-build:
	docker build -t $(USR_NAME)/$(PKG_NAME) .

docker-run:
	docker rm $(PKG_NAME) || true
	docker run --name $(PKG_NAME) \
	-e "TZ=Asia/Novosibirsk" \
	-v ~/.dockerdata/$(PKG_NAME):/data/$(PKG_NAME) \
	$(USR_NAME)/$(PKG_NAME)

clean:
	rm tmp/$(PKG_NAME) || true
	docker rmi -f $(USR_NAME)/$(PKG_NAME)

dev: docker-build docker-run