# this will start the application
# it will use the binary that is built with the build target
start:
	@./bin/main

# this will build the application
# it will compile the go code and create a binary called main
# in the bin directory
build:
	@go build -o ./bin/main cmd/main.go
