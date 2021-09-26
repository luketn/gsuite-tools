.PHONY: build clean deploy gomodgen zip

build: gomodgen
	mkdir -p bin
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap src/go/main.go
	cp -r src/resources/* bin/
	"$(MAKE)" -C bin -f $(CURDIR)/Makefile zip

clean:
	rm -rf bootstrap.zip ./bin ./vendor

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

#Called from in the bin directory
zip:
	zip ../bootstrap.zip $(wildcard *)
