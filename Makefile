ifndef PKGS
PKGS := $(shell go list ./... 2>&1 | grep -v 'github.com/libopenstorage/autopilot-api/vendor' | grep -v 'pkg/client/informers/externalversions' | grep -v versioned | grep -v 'pkg/apis/autopilot')
endif

GO_FILES := $(shell find . -name '*.go' | grep -v vendor | \
                                   grep -v '\.pb\.go' | \
                                   grep -v '\.pb\.gw\.go' | \
                                   grep -v 'externalversions' | \
                                   grep -v 'versioned' | \
                                   grep -v 'generated')

ifeq ($(BUILD_TYPE),debug)
BUILDFLAGS += -gcflags "-N -l"
endif


LDFLAGS += "-s -w"
BUILD_OPTIONS := -ldflags=$(LDFLAGS)

.DEFAULT_GOAL=all
.PHONY: clean vendor vendor-update

all: check-fmt vet lint staticcheck

vendor-update:
	dep ensure -update

vendor:
	dep ensure

check-fmt:
	bash -c "diff -u <(echo -n) <(gofmt -l -d -s -e $(GO_FILES))"


lint: $(GOPATH)/bin/golint
	# golint check ...
	@for file in $(GO_FILES); do $(GOPATH)/bin/golint -set_exit_status $${file} || exit 1; done

vet:
	go vet ./...

$(GOPATH)/bin/golint:
	go install golang.org/x/lint/golint@latest

$(GOPATH)/bin/staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

$(GOPATH)/bin/errcheck:
	go install github.com/kisielk/errcheck@latest


staticcheck: $(GOPATH)/bin/staticcheck
	$(GOPATH)/bin/staticcheck ./...


errcheck: $(GOPATH)/bin/errcheck
	# errcheck check ...
	@$(GOPATH)/bin/errcheck -verbose -blank ./...

pretest: lint vet errcheck staticcheck

codegen:
	@echo "Generating CRD"
	@hack/update-codegen.sh

clean:
	go clean -i $(PKGS)
