#!/bin/bash


exec protoc -I=. --go_out=plugins=grpc:. ./student.proto