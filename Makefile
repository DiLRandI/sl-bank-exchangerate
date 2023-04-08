
build-plugins:
	go build -buildmode=plugin -o ./bin/boc.so plugins/boc/boc.go

build-app:
	go build -o ./bin/slexchangerate ./cmd/slexchangerate/main.go

copy-config:
	cp ./config.json bin

build: build-app build-plugins copy-config

clean: 
	rm -rf ./bin