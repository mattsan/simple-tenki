pkg:
	GOPATH=`pwd` go get github.com/aws/aws-lambda-go/lambda

build: pkg
	GOPATH=`pwd` GOOS=linux go build -o bin/main

deploy: build
	sls deploy

remove:
	sls remove