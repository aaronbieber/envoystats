build/envoystats: envoystats.go
	CGO_ENABLED=0 go build -o build/

.PHONY: build
build: build/envoystats
