all: server css

install_deps:
	go get -v -u github.com/valyala/quicktemplate/qtc
	npm update

server:
	go generate ./...
	go build -v

css:
	$(MAKE) -C less
