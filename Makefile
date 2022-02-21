all: build
.PHONY: all

FFI_PATH:=extern/filecoin-ffi/
FFI_DEPS:=.install-filcrypto
FFI_DEPS:=$(addprefix $(FFI_PATH),$(FFI_DEPS))

$(FFI_DEPS): build-dep/.filecoin-install ;
MODULES:=

CLEAN:=

ldflags=-X=github.com/ipfs-force-community/chain-co/version/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	    ldflags+=-extldflags=$(LDFLAGS)
	endif

GOFLAGS+=-ldflags="$(ldflags)"

build-dep:
	mkdir $@

build-dep/.filecoin-install: $(FFI_PATH) build-dep
	    $(MAKE) -C $(FFI_PATH) $(FFI_DEPS:$(FFI_PATH)%=%)
		    @touch $@

MODULES+=$(FFI_PATH)
BUILD_DEPS+=build-dep/.filecoin-install
CLEAN+=build-dep/.filecoin-install

$(MODULES): build-dep/.update-modules ;

# dummy file that marks the last time modules were updated
build-dep/.update-modules: build-dep
	git submodule update --init --recursive
	touch $@

CLEAN+=build-dep/.update-modules

test: $(BUILD_DEPS)
	go test -v -failfast ./...

lint: $(BUILD_DEPS)
	golint --set_exit_status `go list ./... | grep -v /extern/`

dep-check: build-dep/.update-modules
	./tool/scripts/submodule-check.sh

dist-clean:
	git clean -xdff
	git submodule deinit --all -f

build: $(BUILD_DEPS)
	rm -f chain-co
	go build -o ./chain-co $(GOFLAGS) ./chain-ro/cmd