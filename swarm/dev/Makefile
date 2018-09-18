.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/AERUMTechnology/go-aerum/cmd/swarm
	go build -o bin/geth     github.com/AERUMTechnology/go-aerum/cmd/geth
	go build -o bin/bootnode github.com/AERUMTechnology/go-aerum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/AERUMTechnology/go-aerum/swarm/...
