#!/bin/bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/aws_sqs/main.go
zip bot.zip bootstrap