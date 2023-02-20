win:
	go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64 && go build -ldflags '-extldflags "-static -logg"' -o bin/ConsoleGpt.exe ./cmd/main.go
mac:
	go env -w CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 && go build -ldflags '-extldflags "-static -logg"' -o bin/ConsoleGpt_darwin ./cmd/main.go
linux:
	go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags '-extldflags "-static -logg"' -o bin/ConsoleGpt_linux ./cmd/main.go
all:
	make win && make mac && make linux
clean:
	rm -rf bin/*
