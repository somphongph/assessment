test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test:
	go test -v ./...

test-ing:
	go test -v ./... -tags integration 