# Weather MCP Tool

A Go-based weather query MCP tool that provides global weather information using the OpenWeatherMap API.

## Features

- Global city weather query support
- Natural language query support
- Chinese/English interface switching
- Multi-day weather forecast (1-5 days)
- Temperature unit switching (Celsius/Fahrenheit)
- MCP protocol integration

## Installation

1. Ensure Go 1.21 or higher is installed
2. Clone the repository:
   ```bash
   git clone https://github.com/bytectlgo/mcp_tool.git
   cd mcp_tool
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Configuration

Before using, set up your OpenWeatherMap API key:

```bash
export OPENWEATHERMAP_API_KEY="your_api_key"
```

## Usage

### Using Makefile

1. Build and run:
   ```bash
   # Build all commands
   make build

   # Run weather service
   make run-weather_mcp
   ```

### Using Go Commands

1. Build:
   ```bash
   go build -o bin/weather_mcp ./cmd/weather_mcp
   ```

2. Run:
   ```bash
   ./bin/weather_mcp
   ```

3. Use MCP client to send queries in the following formats:
   ```
   weather query Beijing
   weather query Shanghai weather
   weather query London weather for next 3 days
   weather query New York weather tomorrow
   weather query Paris weather for next 5 days fahrenheit
   weather query Tokyo weather english
   ```

## Query Parameters

| Parameter | Description | Default |
|-----------|-------------|---------|
| city | City name (Chinese/English) | Required |
| days | Forecast days (1-5) | 1 |
| units | Temperature unit (metric: Celsius, imperial: Fahrenheit) | metric |
| lang | Response language (zh: Chinese, en: English) | zh |

## Response Format

```json
{
    "city": "Beijing",
    "temperature": 25.5,
    "description": "Clear",
    "humidity": 45,
    "forecast": [
        {
            "date": "2024-04-24 12:00:00",
            "temperature": 26.0,
            "description": "Cloudy"
        },
        {
            "date": "2024-04-25 12:00:00",
            "temperature": 24.5,
            "description": "Light Rain"
        }
    ]
}
```

## License

MIT License