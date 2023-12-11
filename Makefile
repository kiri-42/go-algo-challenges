DOCKER-EXEC=docker compose exec golang
TARGET=

fix-fmt:
	$(DOCKER-EXEC) gofmt -l -w /go/src/$(TARGET)

fmt:
	! $(DOCKER-EXEC) gofmt -l /go/src/$(TARGET) | tee /dev/tty | read

lint:
	$(DOCKER-EXEC) go vet /go/src/$(TARGET)

test:
	$(DOCKER-EXEC) go test -count=1 /go/src/$(TARGET)

gpt-test:
	$(DOCKER-EXEC) go test -count=1 -run 'Test.*ByChatGPT' /go/src/$(TARGET)

check:
	make fmt
	make lint
	make gpt-test
