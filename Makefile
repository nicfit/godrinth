.DEFAULT_GOAL := build
GODRINTH ?= ./build/godrinth

all: clean build test

build: enums
	go build -o build/godrinth ./cmd/godrinth.go
.PHONY: build

clean:
	rm -rf build/godrinth

clean-dist: clean
	rm -rf build/

test:
	go test ./...
	$(GODRINTH) search sodium
	$(GODRINTH) get lithium
	$(GODRINTH) get gvQqBUqZ

LOCALBIN ?= $(shell pwd)/build/tools
## Location to install dependencies to
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

GO_ENUM_VERSION ?= v0.6.1
GOENUM ?= $(LOCALBIN)/go-enum-$(GO_ENUM_VERSION)
GO_ENUM_URL = https://github.com/abice/go-enum/releases/download/$(GO_ENUM_VERSION)/go-enum_$(shell uname -s)_$(shell uname -m)
$(ENUM_AUTOGEN): GO_ENUM_FLAGS=--marshal --names --ptr
ENUM_AUTOGEN = pkg/facets_enum.go

# The generator statement for go enum files.  Files that invalidate the
# enum file: source file, the binary itself, and this file (in case you want to generate with different flags)
enums: go-enum $(ENUM_AUTOGEN)

%_enum.go: go-enum %.go Makefile
	$(GOENUM) -f $*.go $(GO_ENUM_FLAGS)

# A helper rule for devs to install go-enum
.PHONY: go-enum
go-enum: $(GOENUM)
$(GOENUM): $(LOCALBIN)
	$(call go-install-tool,$(GOENUM),github.com/abice/go-enum,$(GO_ENUM_VERSION))

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef