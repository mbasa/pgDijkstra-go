#!/bin/sh
export CGO_CFLAGS="-I`pg_config --includedir-server` "
go get  github.com/RyanCarrier/dijkstra
go build -buildmode=c-shared -o libPgDijkstra-go.so