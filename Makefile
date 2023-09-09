PROTOS_DIR = proto
PROTOS = \
	$(PROTOS_DIR)/uuid.proto
GO_DIR = go
GO_PACKAGE = github.com/nicholasngai/uuidpb/go/uuidpb
GO_PACKAGE_NAME = uuidpb

.PHONY: all
all: protogen-go

.PHONY: protogen-go
protogen-go:
	rm -rf $(GO_DIR)/gen
	mkdir -p $(GO_DIR)/gen
	protoc \
		--proto_path=$(PROTOS_DIR) \
		--go_out=$(GO_DIR)/gen \
		--go-grpc_out=$(GO_DIR)/gen \
		$(PROTOS)
	@if find $(GO_DIR)/gen -type f -not -path '$(GO_DIR)/gen/$(GO_PACKAGE)/*' | grep ''; then \
		echo 'Found generated stubs outside of $(GO_DIR)/gen/$(GO_PACKAGE)! Please ensure that all protos are defined with go_package within $(GO_PACKAGE).' \
		; exit 1 \
	; fi
	find $(GO_DIR)/$(GO_PACKAGE_NAME) -type f -name '*.pb.go' -print0 | xargs -0 rm -f
	mv $(GO_DIR)/gen/$(GO_PACKAGE)/* $(GO_DIR)/$(GO_PACKAGE_NAME)
	rm -rf $(GO_DIR)/gen
