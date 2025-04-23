package tests

import (
	"os"
	"testing"

	"github.com/briandowns/openweathermap"
)

func TestWeatherAPI(t *testing.T) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		t.Fatal("请设置 API Key")
	}

	// 测试当前天气
	t.Run("测试当前天气", func(t *testing.T) {
		w, err := openweathermap.NewCurrent("C", "en", apiKey)
		if err != nil {
			t.Fatalf("初始化天气客户端失败: %v", err)
		}

		err = w.CurrentByName("Beijing")
		if err != nil {
			t.Fatalf("获取天气信息失败: %v", err)
		}

		t.Logf("城市: %s", w.Name)
		t.Logf("温度: %.1f°C", w.Main.Temp)
		t.Logf("天气描述: %s", w.Weather[0].Description)
		t.Logf("湿度: %d%%", w.Main.Humidity)
	})

	// 测试天气预报
	t.Run("测试天气预报", func(t *testing.T) {
		// 先获取城市的坐标
		w, err := openweathermap.NewCurrent("C", "en", apiKey)
		if err != nil {
			t.Fatalf("初始化天气客户端失败: %v", err)
		}

		err = w.CurrentByName("Beijing")
		if err != nil {
			t.Fatalf("获取天气信息失败: %v", err)
		}

		// 使用坐标获取天气预报
		oc, err := openweathermap.NewOneCall("C", "en", apiKey, nil)
		if err != nil {
			t.Fatalf("初始化 OneCall 客户端失败: %v", err)
		}

		err = oc.OneCallByCoordinates(
			&openweathermap.Coordinates{
				Latitude:  w.GeoPos.Latitude,
				Longitude: w.GeoPos.Longitude,
			},
		)
		if err != nil {
			t.Fatalf("获取预报信息失败: %v", err)
		}

		// 打印预报信息
		for i, day := range oc.Daily {
			if i >= 3 {
				break
			}
			t.Logf("日期: %d", day.Dt)
			t.Logf("温度: %.1f°C - %.1f°C", day.Temp.Min, day.Temp.Max)
			t.Logf("天气描述: %s", day.Weather[0].Description)
			t.Logf("湿度: %d%%", day.Humidity)
			t.Log("-------------------")
		}
	})
}
