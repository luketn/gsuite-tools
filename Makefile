.PHONY: build clean deploy gomodgen zip

SOURCE_FILES=src/go/main.go src/go/staticContent.go

build: gomodgen
	mkdir -p bin
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap $(SOURCE_FILES)
	cp -r src/resources/* bin/
	"$(MAKE)" -C bin -f $(CURDIR)/Makefile zip

build-ui:
	cd src/ui
	npm run build

clean:
	rm -rf bootstrap.zip ./bin ./vendor

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

#Called from in the bin directory
zip:
	zip -r ../bootstrap.zip $(wildcard *)
