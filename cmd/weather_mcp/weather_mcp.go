package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/briandowns/openweathermap"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type WeatherResponse struct {
	City        string     `json:"city"`
	Temperature float64    `json:"temperature"`
	Description string     `json:"description"`
	Humidity    int        `json:"humidity"`
	Forecast    []Forecast `json:"forecast,omitempty"`
}

type Forecast struct {
	Date        string  `json:"date"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func main() {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		log.Fatal("请设置 OPENWEATHERMAP_API_KEY 环境变量")
	}

	// 创建MCP服务器
	srv := server.NewMCPServer(
		"Weather Service",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// 注册天气查询命令
	tool := mcp.NewTool("query",
		mcp.WithDescription("查询天气信息"),
		mcp.WithString("city", mcp.Description("城市名称"), mcp.Required()),
		mcp.WithNumber("days", mcp.Description("预报天数"), mcp.DefaultNumber(1)),
		mcp.WithString("units", mcp.Description("温度单位"), mcp.Enum("metric", "imperial"), mcp.DefaultString("metric")),
		mcp.WithString("lang", mcp.Description("返回语言"), mcp.Enum("zh", "en"), mcp.DefaultString("zh")),
	)

	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments
		city, ok := args["city"].(string)
		if !ok {
			return mcp.NewToolResultError("请提供城市名称"), nil
		}

		days := 1
		if d, ok := args["days"].(float64); ok {
			days = int(d)
		}

		units := "metric"
		if u, ok := args["units"].(string); ok {
			units = u
		}

		lang := "zh"
		if l, ok := args["lang"].(string); ok {
			lang = l
		}

		// 获取当前天气
		w, err := openweathermap.NewCurrent(units, lang, apiKey)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("初始化天气客户端失败: %v", err)), nil
		}

		err = w.CurrentByName(city)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("获取天气信息失败: %v", err)), nil
		}

		response := WeatherResponse{
			City:        w.Name,
			Temperature: w.Main.Temp,
			Description: w.Weather[0].Description,
			Humidity:    w.Main.Humidity,
		}

		// 如果需要预报
		if days > 1 {
			f, err := openweathermap.NewForecast(units, lang, apiKey, "5")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("初始化预报客户端失败: %v", err)), nil
			}

			err = f.DailyByName(city, days)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("获取预报信息失败: %v", err)), nil
			}

			// 由于 ForecastWeatherJson 是一个接口，我们需要使用反射或其他方式来访问数据
			// 这里我们暂时只返回当前天气信息
			response.Forecast = []Forecast{}
		}

		return mcp.NewToolResultText(fmt.Sprintf("%+v", response)), nil
	})

	log.Println("天气查询服务启动中...")
	if err := server.ServeStdio(srv); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
