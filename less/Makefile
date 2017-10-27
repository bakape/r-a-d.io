all: $(addsuffix .css, $(addprefix ../www/css/, $(basename $(wildcard *.less))))

../www/css/%.css: %.less
	../node_modules/.bin/lessc $^ | ../node_modules/.bin/cleancss -o $@
