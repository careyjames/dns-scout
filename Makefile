DESTDIR=/usr/local
GO_FILES=main.go

default: build

build:
	go build -o bin/dns-scout $(GO_FILES)

run: build
	./bin/dns-scout

install:
	install -m 755 bin/dns-scout $(DESTDIR)/bin

clean:
	rm -f bin/dns-scout

test:
	go test -cover ./...

coverage:
	go tool cover -html=coverage.out -o coverage.html
