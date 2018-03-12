clean:
	rm ./price
build:
	make clean
	go build
test:
	make build
	./price one two three four