all: server client css

.PHONY: client

install_deps:
	go get -v -u github.com/valyala/quicktemplate/qtc
	npm update

server:
	rm -f templates/*.qtpl.go
	go generate ./...
	go build -v

client:
	mkdir -p www/js
	node_modules/.bin/tsc -p client --outdir www/js
	$(foreach i, $(wildcard www/js/*), node_modules/.bin/uglifyjs -o $i $i)

css:
	$(MAKE) -C less
