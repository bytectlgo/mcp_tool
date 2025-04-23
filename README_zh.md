# 天气 MCP 工具

一个简洁的天气查询工具，让你用一句话就能查询全球天气，完美集成 Cursor 编辑器。

## ✨ 功能特点

- 💡 **简洁**: 一行命令查询天气
- 🤖 **智能**: 支持中英文自然语言
- 🌏 **全球**: 支持所有主要城市
- 🔌 **即插即用**: 完美集成 Cursor
- 🚀 **高性能**: 异步处理，快速响应
- 🎨 **美观**: 清晰直观的天气显示

## 🚀 快速开始

### 1. 获取 API Key

> 🔑 开始前，请先获取 OpenWeather API Key

### 2. 安装

#### 2.1 使用 Homebrew 安装（推荐）

```bash
# 添加 tap 源
brew tap bytectlgo/homebrew-tap

# 安装工具
brew install mcp_tool
```

#### 2.2 手动安装

##### 2.2.1 克隆并安装

```bash
git clone https://github.com/bytectlgo/mcp_tool.git
cd mcp_tool
go mod download
```

##### 2.2.2 配置 API Key

**方法1：使用配置文件（推荐）**

复制示例配置文件并修改：

```bash
cp env.example .env
```

然后编辑 `.env` 文件，将 `your_api_key_here` 替换为你的 API Key。

**方法2：使用环境变量**

macOS/Linux:
```bash
export OPENWEATHERMAP_API_KEY="your_api_key"
```

Windows:
```bash
set OPENWEATHERMAP_API_KEY=your_api_key
```

##### 2.2.3 编译和运行

使用 Makefile 编译和运行：

```bash
# 编译所有命令
make build

# 运行天气查询服务
make run-weather_mcp
```

或者直接使用 Go 命令：

```bash
# 编译
go build -o bin/weather_mcp ./cmd/weather_mcp

# 运行
./bin/weather_mcp
```

#### 2.3 启用工具

编辑 `~/.cursor/mcp.json`（Windows：`%USERPROFILE%\.cursor\mcp.json`）：

```json
{
    "weather_mcp": {
        "command": "weather_mcp"
    }
}
```

重启 Cursor 即可使用！

## 📝 使用示例

在 Cursor 中直接输入：

```
北京的天气怎么样？
上海未来3天天气
伦敦明天会下雨吗？
纽约的天气 fahrenheit
东京天气 english
```

## ⚙️ 参数说明

| 参数 | 说明 | 默认值 |
|------|------|--------|
| city | 城市名称（中英文） | 必填 |
| days | 预报天数（1-5） | 1 |
| units | 温度单位（metric: 摄氏度, imperial: 华氏度） | metric |
| lang | 返回语言（zh: 中文, en: 英文） | zh |

## ❓ 常见问题

1. **无法使用？**  
   - 检查 API Key 是否正确设置  
   - 重启 Cursor  
   - 检查 Go 环境
2. **找不到城市？**  
   - 尝试使用英文名称  
   - 检查拼写  
   - 使用完整的城市名称

## 👨‍💻 作者

* 字节控制
* Email: your-email@example.com

## 🙏 致谢

* FastMCP
* OpenWeatherMap
* Cursor

## 📄 许可证

本项目采用 MIT 许可证 - 详见 LICENSE 文件 