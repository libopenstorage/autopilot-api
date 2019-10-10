DOCKER_HUB_AUTOPILOT_TAG ?= latest
AUTOPILOT_IMG=$(DOCKER_HUB_REPO)/autopilot:$(DOCKER_HUB_AUTOPILOT_TAG)

ifndef PKGS
PKGS := $(shell go list ./... 2>&1 | grep -v 'github.com/libopenstorage/autopilot/vendor' | grep -v 'pkg/client/informers/externalversions' | grep -v versioned | grep -v 'pkg/apis/autopilot')
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

RELEASE_VER := 0.7.0
BASE_DIR    := $(shell git rev-parse --show-toplevel)
GIT_SHA     := $(shell git rev-parse --short HEAD)
BIN         :=$(BASE_DIR)/bin

VERSION = $(RELEASE_VER)-$(GIT_SHA)

LDFLAGS += "-s -w -X github.com/libopenstorage/autopilot/pkg/version.Version=$(VERSION)"
BUILD_OPTIONS := -ldflags=$(LDFLAGS)

.DEFAULT_GOAL=all
.PHONY: test clean vendor vendor-update container deploy mockgen

all: autopilot check-fmt vet lint staticcheck

vendor-update:
	dep ensure -update

vendor:
	dep ensure

check-fmt:
	bash -c "diff -u <(echo -n) <(gofmt -l -d -s -e $(GO_FILES))"

lint:
	go get -v golang.org/x/lint/golint
	for file in $(GO_FILES); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

vet:
	go vet $(PKGS)

$(GOPATH)/bin/staticcheck:
	go get -u honnef.co/go/tools/cmd/staticcheck

staticcheck: $(GOPATH)/bin/staticcheck
	$(GOPATH)/bin/staticcheck $(PKGS)

errcheck:
	go get -v github.com/kisielk/errcheck
	errcheck -verbose -blank $(PKGS)

pretest: lint vet errcheck staticcheck

test:
	# start local prometheus server
	-docker rm -f autopilot-prometheus-test
	docker run -d --name autopilot-prometheus-test -p 9090:9090 prom/prometheus:v2.11.1
	sleep 15
	echo "" > coverage.txt
	for pkg in $(PKGS);	do \
		go test -v -tags unittest -coverprofile=profile.out -covermode=atomic $(BUILD_OPTIONS) $${pkg} || exit 1; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done
	-docker rm -f autopilot-prometheus-test

codegen:
	@echo "Generating CRD"
	@hack/update-codegen.sh

autopilot:
	@echo "Building the autopilot binary"
	@cd cmd/autopilot && CGO_ENABLED=0 go build $(BUILD_OPTIONS) -o $(BIN)/autopilot

container: autopilot
	@echo "Building container: docker build --tag $(AUTOPILOT_IMG) -f Dockerfile ."
	sudo docker build --tag $(AUTOPILOT_IMG) -f Dockerfile .

deploy: container
	sudo docker push $(AUTOPILOT_IMG)

mockgen: $(GOPATH)/bin/mockgen
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen
	mockgen -destination=pkg/mock/storage/driver.mock.go  -package=storage github.com/libopenstorage/autopilot/drivers/storage Driver
	mockgen -destination=pkg/mock/scheduler/driver.mock.go  -package=scheduler github.com/libopenstorage/autopilot/drivers/scheduler Driver
	mockgen -destination=pkg/mock/monitoring/provider.mock.go  -package=monitoring github.com/libopenstorage/autopilot/drivers/monitoring Provider
	mockgen -destination=pkg/mock/rule/rule.mock.go  -package=rule github.com/libopenstorage/autopilot/pkg/rule Rule

clean:
	go clean -i $(PKGS)
