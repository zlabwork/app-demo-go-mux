#!/bin/bash

export $(cat .env | grep -v "#")
go run cmd/main.go
