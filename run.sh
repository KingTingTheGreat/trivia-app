#!/bin/bash

go run setup/setup.go
go build -o trivia-app.exe
./trivia-app.exe

