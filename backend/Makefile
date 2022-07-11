BUILD_DIR=tmp
SUB_DIRS=pipeline
BUILD_TARGETS=$(addprefix cmd/,$(SUB_DIRS))

all: deps $(BUILD_TARGETS)
$(BUILD_TARGETS): %: 
	go build -o '$(BUILD_DIR)/$(subst cmd-,$e,$(subst /,-,$@))' '$@'/main.go
.PHONY: deps $(BUILD_TARGETS)

deps: 
	go mod tidy
	go get ./...

clean:
	rm build/*