.PHONY: build

build:
	@if [ ! -d "build" ]; then mkdir "build"; fi

	@if [ ! -d "build/windows" ]; then mkdir "build/windows"; fi
	@if [ ! -d "build/windows/amd64" ]; then mkdir "build/windows/amd64"; fi
	GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/boom.exe ./cmd/cli

	@if [ ! -d "build/linux" ]; then mkdir "build/linux"; fi
	@if [ ! -d "build/linux/amd64" ]; then mkdir "build/linux/amd64"; fi
	GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/boom ./cmd/cli
