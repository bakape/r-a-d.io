all: server client css

.PHONY: client

install_deps:
	go get -v -u github.com/valyala/quicktemplate/qtc
	npm update

server:
	go generate ./...
	go build -v


#TODO: Minification and sourcemaps
client:
	mkdir -p www/js
	node_modules/.bin/tsc -p client --outdir www/js

css:
	$(MAKE) -C less
