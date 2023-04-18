all: build
.PHONY: all

## variables

# git modules that need to be loaded
MODULES:=

ldflags=-X=github.com/ipfs-force-community/chain-co/version.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	    ldflags+=-extldflags=$(LDFLAGS)
	endif

GOFLAGS+=-ldflags="$(ldflags)"

## FFI

FFI_PATH:=extern/filecoin-ffi/
FFI_DEPS:=.install-filcrypto
FFI_DEPS:=$(addprefix $(FFI_PATH),$(FFI_DEPS))

$(FFI_DEPS): build-dep/.filecoin-install ;

build-dep/.filecoin-install: $(FFI_PATH)
	$(MAKE) -C $(FFI_PATH) $(FFI_DEPS:$(FFI_PATH)%=%)
	@touch $@

MODULES+=$(FFI_PATH)
BUILD_DEPS+=build-dep/.filecoin-install
CLEAN+=build-dep/.filecoin-install

## modules
build-dep:
	mkdir $@

$(MODULES): build-dep/.update-modules;
# dummy file that marks the last time modules were updated
build-dep/.update-modules: build-dep;
	git submodule update --init --recursive
	touch $@

## build

test:
	go test -race ./...

lint: $(BUILD_DEPS)
	golangci-lint run

dist-clean:
	git clean -xdff
	git submodule deinit --all -f

build: $(BUILD_DEPS)
	rm -f chain-co
	go build -o ./chain-co $(GOFLAGS) ./cmd

proxy-gen:
	go run ./gen/proxy
.PHONY: proxy-gen

perm-gen:
	go run ./gen/perm
.PHONY: perm-gen

gen-all: proxy-gen perm-gen
	go generate ./...
.PHONY: gen-all


TAG:=test
docker: $(BUILD_DEPS)
ifdef DOCKERFILE
	cp $(DOCKERFILE) ./dockerfile
else
	curl -O https://raw.githubusercontent.com/filecoin-project/venus-docs/master/script/docker/dockerfile
endif
	docker build --build-arg HTTPS_PROXY=$(BUILD_DOCKER_PROXY) --build-arg BUILD_TARGET=chain-co -t chain-co .
	docker tag chain-co filvenus/chain-co:$(TAG)

ifdef PRIVATE_REGISTRY
	docker tag chain-co $(PRIVATE_REGISTRY)/filvenus/chain-co:$(TAG)
endif

docker-push: docker
	docker push $(PRIVATE_REGISTRY)/filvenus/chain-co:$(TAG)
