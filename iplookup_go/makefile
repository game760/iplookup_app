
BINARY_LINUX = ./bin/iplookup_linux
BINARY_WIN = ./bin/iplookup_win.exe
GOARCH ?= amd64

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -o $(BINARY_LINUX) main.go
	@echo "Linux Done: $(BINARY_LINUX)"

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) go build -o $(BINARY_WIN) main.go
	@echo "Windows Done: $(BINARY_WIN)"

linux: build-linux

win: build-windows


clean:
	rm -f $(BINARY_LINUX) $(BINARY_WIN)
	@echo "已清理所有编译文件"


run:
	@if [ "$(shell uname)" = "Linux" ]; then \
		./$(BINARY_LINUX); \
	elif [ "$(shell uname -s | cut -d' ' -f1)" = "MINGW64" ] || [ "$(shell uname -s)" = "Windows_NT" ]; then \
		./$(BINARY_WIN); \
	else \
		echo "不支持的系统"; \
	fi