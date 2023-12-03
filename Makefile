build:
	@go build -o ./bin/gopoker

run: 
	@make build
	@./bin/gopoker

test:
	go test -v ./...

clean:
	@rm -rf bin