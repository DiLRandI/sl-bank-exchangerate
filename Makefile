include version.mk

build-plugins:
	go build -buildmode=plugin -ldflags "-X main.Version=${PLUGIN_BOC}" -o ./bin/boc.so ./cmd/plugins/boc/boc.go
	go build -buildmode=plugin -ldflags "-X main.Version=${PLUGIN_HNB}" -o ./bin/hnb.so ./cmd/plugins/hnb/hnb.go

build-app:
	go build -ldflags "-X main.Version=${APP_VERSION}" -o ./bin/slexchangerate ./cmd/slexchangerate/main.go

copy-config:
	cp ./config.json bin

build: clean build-app build-plugins copy-config

clean: 
	rm -rf ./bin