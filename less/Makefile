all: $(addsuffix .css, $(addprefix ../www/css/, $(basename $(wildcard *.less))))

../www/css/%.css: %.less
	mkdir -p ../www/css
	../node_modules/.bin/lessc --clean-css $^ $@
