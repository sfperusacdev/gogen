VERSION := v0.0.1
build:
	go build -o build/gen-$(VERSION)-linux-amd64 main.go	
	env GOOS=windows GOARCH=amd64 go build -o build/gen-$(VERSION)-amd64.exe main.go	


