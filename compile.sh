#!/bin/sh
export CGO_CFLAGS="-I`pg_config --includedir-server` "
go build -buildmode=c-shared -o libPgDijkstra-go.so

#CGO_CFLAGS="-I/opt/homebrew/Cellar/postgresql@14/14.6/include/postgresql@14/server/" go build -buildmode=c-shared -o libGoPgFunc.so