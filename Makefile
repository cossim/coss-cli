GO := go

# 项目名称
APP_NAME := coss-cli

# 版本号 (可以通过环境变量传入)
VERSION ?= 1.0.4

# 交叉编译目标平台 (可以通过环境变量传入)
PLATFORMS ?= linux windows darwin

# 支持的架构 (可以通过环境变量传入)
ARCHS ?= amd64 arm64

# 构建目录
BUILD_DIR := build

# 构建参数
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

# 默认目标
.DEFAULT_GOAL := build

# 构建目标
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-$(GOOS)-$(GOARCH)-$(VERSION)$(EXT)

# 打包目标
.PHONY: package
package: $(foreach platform,$(PLATFORMS),$(foreach arch,$(ARCHS),package-$(platform)-$(arch)))

package-%: VERSION ?= $(VERSION)
package-%:
	@GOOS=$(word 1,$(subst -, ,$*)) GOARCH=$(word 2,$(subst -, ,$*)) EXT=$(if $(findstring windows,$(word 1,$(subst -, ,$*))),.exe) $(MAKE) build

# 清理目标
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

# 帮助目标
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build       Build the Go program for the current platform"
	@echo "  package     Package the Go program for all supported platforms"
	@echo "  clean       Clean the build directory"
	@echo "  release     Package and release the Go program with a specific tag and version"

# 发布目标
.PHONY: release
release: VERSION ?= $(VERSION)
release: TAG ?= $(TAG)
release: package
	@gh release create $(TAG) $(BUILD_DIR)/*-$(VERSION)* --title "$(TAG)" --notes "Release version $(VERSION)"

# 示例用法
.PHONY: example
example:
	@echo "Example: make release VERSION=1.0.2 TAG=v1.0.2"
