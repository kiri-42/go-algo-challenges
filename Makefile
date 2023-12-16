TARGET=
DOCKER-EXEC=docker compose exec -w /go/src/$(TARGET) golang

init:
	mkdir $(TARGET)
	$(DOCKER-EXEC) go mod init github.com/kiri-42/go-algo-challenges/$(TARGET)
	docker compose exec golang go work use $(TARGET)
	touch $(TARGET)/question.md $(TARGET)/submit.go $(TARGET)/submit_test.go

solution:
	touch $(TARGET)/evaluation_test.go $(TARGET)/solution.md

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
