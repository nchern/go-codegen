.PHONY: clean install install-deps build test

TEST_DIR=tests

GEN=go-codegen --pkg=main --sync 


install-deps:
	go get -u github.com/jteeuwen/go-bindata/...

clean:
	rm -rf $(TEST_DIR)/generated_*.go

install:
	go-bindata -pkg=code -prefix=code -o code/bindata.go code/templates/
	go get ./...

build:
	go-bindata -pkg=code -prefix=code -o code/bindata.go code/templates/
	go build ./...

test:
	$(GEN) --name=StringMap map string string | goimports > $(TEST_DIR)/generated_string_map.go
	$(GEN) --name=StringList list string | goimports > $(TEST_DIR)/generated_string_list.go

	go test -race ./...
