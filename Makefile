exec=mehmet

.DEFAULT_GOAL := build

configure: clean
	
test: configure
	go vet ./...
	go test ./...
build: configure
	go build -v  -o ${exec}
build-linux: configure
	GOOS=linux GOARCH=amd64 go build -v  -o ${exec}
clean:
	rm -f ${exec}
