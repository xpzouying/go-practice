#!/bin/bash


exec protoc -I=. --go_out=. ./student.proto