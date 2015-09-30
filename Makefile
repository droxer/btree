default: deps test cover

deps:
	go get -d -v ./...
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs -n1 go get -d

fmt:
	gofmt -w .

test:
	go test ./...

cover:
	go test ./... --cover

.PHONY: deps test cover fmt