#!/bin/bash

go run setup/setup.go
echo "Compiling program"
go build -o trivia-app.exe
echo "Running program"
./trivia-app.exe

