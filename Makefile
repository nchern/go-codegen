.PHONY: clean install install-deps build test bindata

TEST_DIR=tests

GEN=go-codegen --pkg=main --sync 


install-deps:
	go get -u github.com/jteeuwen/go-bindata/...

clean:
	rm -rf $(TEST_DIR)/generated_*.go

bindata:
	go-bindata -pkg=code -prefix=code -o code/bindata.go code/templates/

install: bindata
	go get ./...

build: bindata
	go build ./...

test:
	$(GEN) --name=StringMap map string string | goimports > $(TEST_DIR)/generated_string_map.go
	$(GEN) --name=StringList list string | goimports > $(TEST_DIR)/generated_string_list.go

	go test -race ./...
