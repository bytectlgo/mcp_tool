# 定义变量
GO=go
GOFLAGS=-v
LDFLAGS=-ldflags="-s -w"
BIN_DIR=bin

# 获取所有命令
CMDS=$(shell ls -d cmd/*)

# 默认目标
all: build

# 创建 bin 目录
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

# 编译所有命令
build: $(BIN_DIR)
	@echo "正在编译所有命令..."
	@for cmd in $(CMDS); do \
		cmd_name=$$(basename $$cmd); \
		echo "编译 $$cmd_name..."; \
		$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BIN_DIR)/$$cmd_name ./cmd/$$cmd_name; \
	done

# 清理编译产物
clean:
	@echo "清理编译产物..."
	@rm -rf $(BIN_DIR)

# 运行指定命令
run-%: build
	@echo "运行 $*..."
	@./$(BIN_DIR)/$*

# 安装所有命令
install: build
	@echo "安装所有命令..."
	@for cmd in $(CMDS); do \
		cmd_name=$$(basename $$cmd); \
		echo "安装 $$cmd_name..."; \
		cp $(BIN_DIR)/$$cmd_name /usr/local/bin/; \
	done

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make build           - 编译所有命令"
	@echo "  make clean           - 清理编译产物"
	@echo "  make run-<cmd>       - 编译并运行指定命令"
	@echo "  make install         - 安装所有命令到系统"
	@echo "  make help            - 显示帮助信息"
	@echo ""
	@echo "可用命令列表:"
	@for cmd in $(CMDS); do \
		cmd_name=$$(basename $$cmd); \
		echo "  - $$cmd_name"; \
	done

.PHONY: all build clean run-% install help 