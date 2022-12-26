test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test:
	go test -v ./...

test-ingt:
	AUTH_TOKEN="Basic YXBpZGVzaWduOjQ1Njc4" go test -v ./... -tags=integration 