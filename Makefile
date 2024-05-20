# Go 编译器
GO := go

# 项目名称
APP_NAME := coss-cli

# 版本号
VERSION := 1.0.1

# 交叉编译目标平台
PLATFORMS := linux windows darwin

# 支持的架构
ARCHS := amd64

# 构建目录
BUILD_DIR := build

# 构建参数
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

# 默认目标
.DEFAULT_GOAL := build

# 构建目标
.PHONY: build
build:
	@$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-$(VERSION)$(EXT)

# 打包目标
.PHONY: package
package: $(foreach platform,$(PLATFORMS),$(foreach arch,$(ARCHS),$(platform)-$(arch)))

$(foreach platform,$(PLATFORMS),$(foreach arch,$(ARCHS),$(platform)-$(arch))):
	@if [ "$(GOOS)" != "$(word 1,$(subst -, ,$@))" ] || [ "$(GOARCH)" != "$(word 2,$(subst -, ,$@))" ]; then \
		$(MAKE) build GOOS=$(word 1,$(subst -, ,$@)) GOARCH=$(word 2,$(subst -, ,$@)) EXT=$(if $(findstring windows,$(word 1,$(subst -, ,$@))),.exe); \
	fi

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
	@echo "  build     Build the Go program for the current platform"
	@echo "  package   Package the Go program for all supported platforms"
	@echo "  clean     Clean the build directory"