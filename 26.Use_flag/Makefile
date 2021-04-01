build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o release/helloword

docker:
	docker build -t arnold/helloword .

