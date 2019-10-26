TEST_DIR=tests

GEN=go-codegen

.PHONY: clean
clean:
	@find $(TEST_DIR) -name 'generated_*.go' -exec rm {} \;

.PHONY: bindata
bindata:
	go-bindata -prefix=pkg/generic -ignore="_test.go" -pkg=generic -o pkg/generic/bindata.go pkg/generic/list/ pkg/generic/set/ pkg/generic/typedmap/ pkg/generic/iterator/

.PHONY: build
build: bindata vet gen
	go build ./...

.PHONY: install
install: test
	go get ./...

.PHONY: vet
vet:
	 go vet ./...

.PHONY: gen
gen:
	$(GEN) --pkg=main generic -f pkg/generic/typedmap/typedmap.go string string | goimports > $(TEST_DIR)/generic/generated_string_map.go
	$(GEN) --pkg=main generic -f pkg/generic/list/list.go string | goimports > $(TEST_DIR)/generic/generated_string_list.go
	$(GEN) --pkg=main generic -f pkg/generic/iterator/iterator.go  string | goimports | gofmt > $(TEST_DIR)/generic/generated_string_iterator.go
	$(GEN) --pkg=model immutable -f $(TEST_DIR)/immutable/model/model.go | gofmt  > $(TEST_DIR)/immutable/model/generated_model_impl.go

.PHONY: test
test: clean gen
	go test -race ./...
