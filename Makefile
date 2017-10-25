APP = main
MAIN = main.go
build:
	go build -o $(APP) $(MAIN)
update:
	go get -u ./...
package: build
	tar -cvf $(APP).tar.gz $(APP) config.json init.sh Makefile
install:
	mv -f $(APP) $(PREFIX)/bin
fmt:
	go fmt ./...
lint:
	go vet ./...
test:
	go test -short ./...
test-all: lint
	go test ./...
run-debug:
	./$(APP)
run:
	nohup ./$(APP) > debug.log 2>&1 &
clean:
	rm -f $(APP)
	rm -f $(APP).exe
	rm -f $(PREFIX)/bin/$(APP)
.PHONY: build update package install fmt lint test test-all clean all 
