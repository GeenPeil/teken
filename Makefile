
.PHONY: build
build: build-pechtold


.PHONY: build-pechtold
build-pechtold:
	go build -o dist/bin/pechtold github.com/GeenPeil/teken/pechtold
