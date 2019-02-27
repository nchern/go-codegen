TEST_DIR=tests

GEN=go-codegen --pkg=main 

.PHONY: clean
clean:
	rm -rf $(TEST_DIR)/generated_*.go

.PHONY: bindata
bindata:
	go-bindata -prefix=pkg/generic -ignore="_test.go" -pkg=generic -o pkg/generic/bindata.go pkg/generic/list/ pkg/generic/set/ pkg/generic/typedmap/ pkg/generic/iterator/

.PHONY: build
build: bindata
	go build ./...

.PHONY: install
install: build
	go get ./...

.PHONY: test
test: clean
	$(GEN) generic -f pkg/generic/typedmap/typedmap.go string string | goimports > $(TEST_DIR)/generic/generated_string_map.go
	$(GEN) generic -f pkg/generic/list/list.go string | goimports > $(TEST_DIR)/generic/generated_string_list.go
	$(GEN) --pkg=main generic -f pkg/generic/iterator/iterator.go  string | goimports | gofmt > $(TEST_DIR)/generic/generated_string_iterator.go
	$(GEN) --pkg=model immutable -f $(TEST_DIR)/immutable/model/model.go | gofmt  > $(TEST_DIR)/immutable/model/generated_model_impl.go

	go test -race ./...
