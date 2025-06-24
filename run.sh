#!/bin/bash

touch .env
go run setup/setup.go
echo "Compiling program"
go build -o trivia-app.exe ./cmd/main.go
echo "Running program"
./trivia-app.exe

