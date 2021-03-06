GEN=go-codegen
TEST_DIR=tests


.PHONY: bindata
bindata:
	go-bindata -prefix=pkg/generic -ignore="_test.go" -pkg=generic -o\
		pkg/generic/bindata.go\
		pkg/generic/list/\
		pkg/generic/set/\
		pkg/generic/hashmap/\
		pkg/generic/iterator/\
		pkg/generic/converter

.PHONY: build
build: bindata vet
	go build ./...

.PHONY: install
install: test vet bindata
	go install ./...

.PHONY: vet
vet:
	 go vet ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: gen
gen:
	$(GEN) --pkg=main generic -f pkg/generic/hashmap/hashmap.go string string | goimports > $(TEST_DIR)/generic/generated_string_map.go
	$(GEN) --pkg=main generic -f pkg/generic/list/list.go string | goimports > $(TEST_DIR)/generic/generated_string_list.go
	$(GEN) --pkg=main generic -f pkg/generic/iterator/iterator.go string | goimports | gofmt > $(TEST_DIR)/generic/generated_string_iterator.go
	$(GEN) --pkg=model immutable -f $(TEST_DIR)/immutable/model/model.go | gofmt  > $(TEST_DIR)/immutable/model/generated_model_impl.go

.PHONY: test
test: vet
	go test -race ./...
