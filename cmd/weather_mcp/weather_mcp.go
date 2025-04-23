package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	Date        int64   `json:"date"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
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
		mcp.WithString("units", mcp.Description("温度单位"), mcp.Enum("C", "F"), mcp.DefaultString("C")),
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

		units := "C"
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
			// 使用 OneCall API 获取天气预报
			oc, err := openweathermap.NewOneCall(units, lang, apiKey, nil)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("初始化预报客户端失败: %v", err)), nil
			}

			err = oc.OneCallByCoordinates(
				&openweathermap.Coordinates{
					Latitude:  w.GeoPos.Latitude,
					Longitude: w.GeoPos.Longitude,
				},
			)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("获取预报信息失败: %v", err)), nil
			}

			// 添加预报信息
			response.Forecast = make([]Forecast, 0, days)
			for i, day := range oc.Daily {
				if i >= days {
					break
				}
				response.Forecast = append(response.Forecast, Forecast{
					Date:        int64(day.Dt),
					TempMin:     day.Temp.Min,
					TempMax:     day.Temp.Max,
					Description: day.Weather[0].Description,
					Humidity:    day.Humidity,
				})
			}
		}

		// 格式化当前天气信息
		currentWeather := fmt.Sprintf("城市: %s\n当前温度: %.1f°%s\n天气状况: %s\n湿度: %d%%",
			response.City, response.Temperature, units, response.Description, response.Humidity)

		// 如果有预报信息，添加预报
		if len(response.Forecast) > 0 {
			currentWeather += "\n\n未来天气预报:"
			for _, forecast := range response.Forecast {
				date := time.Unix(forecast.Date, 0).Format("2006-01-02")
				currentWeather += fmt.Sprintf("\n%s: %.1f°%s ~ %.1f°%s, %s, 湿度: %d%%",
					date, forecast.TempMin, units, forecast.TempMax, units, forecast.Description, forecast.Humidity)
			}
		}

		return mcp.NewToolResultText(currentWeather), nil
	})

	log.Println("天气查询服务启动中...")
	if err := server.ServeStdio(srv); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
