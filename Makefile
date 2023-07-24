VERSION := v0.0.1
build:
	go build -o build/gen-$(VERSION) main.go	
	env GOOS=windows GOARCH=amd64 go build -o build/gen-$(VERSION).exe main.go	


