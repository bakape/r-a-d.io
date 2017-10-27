all: generate css

install_deps:
	go get -v -u github.com/valyala/quicktemplate/qtc
	npm update

generate:
	go generate ./...

css:
	$(MAKE) -C less
