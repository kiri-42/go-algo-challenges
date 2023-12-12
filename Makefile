TARGET=
DOCKER-EXEC=docker compose exec -w /go/src/$(TARGET) golang

init:
	$(DOCKER-EXEC) go mod init

use:
	docker compose exec golang go work use $(TARGET)

fix-fmt:
	$(DOCKER-EXEC) go fmt

fmt:
	! $(DOCKER-EXEC) gofmt -l -w . | tee /dev/tty | read

lint:
	$(DOCKER-EXEC) go vet

test:
	$(DOCKER-EXEC) go test -count=1

gpt-test:
	$(DOCKER-EXEC) go test -count=1 -run 'Test.*ByChatGPT'

evaluate:
	make fmt
	make lint
	make gpt-test
