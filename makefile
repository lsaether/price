clean:
	touch price
	rm ./price
build:
	make clean
	go build
test:
	make build
	./price -source=coinmarketcap two three four