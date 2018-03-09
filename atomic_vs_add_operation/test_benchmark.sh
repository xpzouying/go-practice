#!/bin/bash

echo "benchkmark mutex"
go test -v ./bm_mutex_test.go ./add_op_with_mutex.go -bench=.


echo "benchkmark atomic"
go test -v ./bm_atomic_test.go ./add_op_with_mutex.go -bench=.

