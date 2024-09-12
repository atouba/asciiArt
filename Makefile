.DEFAULT_GOAL := run

build:
	go build -o ascii_art

run: build
	./ascii_art
