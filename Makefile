
build-plugins:
	go build -buildmode=plugin -o ./bin/boc.so ./cmd/plugins/boc/boc.go
	go build -buildmode=plugin -o ./bin/hnb.so ./cmd/plugins/hnb/hnb.go

build-app:
	go build -o ./bin/slexchangerate ./cmd/slexchangerate/main.go

copy-config:
	cp ./config.json bin

build: build-app build-plugins copy-config

clean: 
	rm -rf ./bin