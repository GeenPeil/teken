
.PHONY: build
build: build-cupido build-kantklosser build-maillist


.PHONY: build-cupido
build-cupido:
	go build -o dist/bin/cupido github.com/GeenPeil/teken/cupido/cmd/cupido

.PHONY: build-kantklosser
build-kantklosser:
	go build -o dist/bin/kantklosser github.com/GeenPeil/teken/kantklosser

.PHONY: build-maillist
build-maillist:
	go build -o dist/bin/maillist github.com/GeenPeil/teken/maillist

