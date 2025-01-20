.PHONY: default run bui9ld test docs clean

# Variabbles

APP_NAME = "goopportubities"

#Tasks 
default: run-with-docs
run: 
	go run main.go

run-with-docs:
	swag init
	go run main.go
build:
	go build -o $(APP_NAME) main.go
test:
	go test -v ./...
docs:
	swag init 
clean:
	rm -rf $(APP_NAME)
	rm -rf docs
