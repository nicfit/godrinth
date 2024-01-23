
build: enums
	go build -o build/godrinth ./cmd/godrinth.go
.PHONY: build

clean:
	rm -rf build/

test:
	go test ./...


GOENUM ?= go-enum
GO_ENUM_VERSION ?= v0.6.0
GO_ENUM_URL = https://github.com/abice/go-enum/releases/download/$(GO_ENUM_VERSION)/go-enum_$(shell uname -s)_$(shell uname -m)
ENUM_AUTOGEN = pkg/facets_enum.go
$(ENUM_AUTOGEN): GO_ENUM_FLAGS=--marshal --names --ptr

# The generator statement for go enum files.  Files that invalidate the
# enum file: source file, the binary itself, and this file (in case you want to generate with different flags)
enums: $(ENUM_AUTOGEN)

%_enum.go: %.go Makefile
	$(GOENUM) -f $*.go $(GO_ENUM_FLAGS)

# A helper rule for devs to install go-enum
go-enum:
	curl -fsSL $(GO_ENUM_URL) -o $(GOENUM)
	chmod 755 $(GOENUM)
.PHONY: go-enum