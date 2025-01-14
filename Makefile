test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test:
	go test -v ./...

test-it:
	AUTH_TOKEN="Basic YXBpZGVzaWduOjQ1Njc4" go test -v -tags=integration  ./... 


docker-build:
	docker build -t my-golang-app .

docker-run:
	docker run -it --rm --name my-running-app -p 2565:2565 --env-file .env my-golang-app

docker-it:
	docker-compose -f docker-compose.integration.yml up --build --abort-on-container-exit --exit-code-from it_tests