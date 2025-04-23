package weather

import (
	"strconv"
	"strings"
)

// QueryParams 表示天气查询参数
type QueryParams struct {
	City  string
	Days  int
	Units string
	Lang  string
}

// ParseQuery 解析天气查询字符串
func ParseQuery(query string) QueryParams {
	// 默认值
	params := QueryParams{
		Days:  1,
		Units: "metric", // 摄氏度
		Lang:  "zh",     // 中文
	}

	// 解析查询字符串
	parts := strings.Fields(query)
	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case "in", "for", "at":
			if i+1 < len(parts) {
				params.City = parts[i+1]
			}
		case "days", "day":
			if i+1 < len(parts) {
				if d, err := strconv.Atoi(parts[i+1]); err == nil {
					params.Days = d
				}
			}
		case "celsius", "c":
			params.Units = "metric"
		case "fahrenheit", "f":
			params.Units = "imperial"
		case "chinese", "zh":
			params.Lang = "zh"
		case "english", "en":
			params.Lang = "en"
		}
	}

	return params
}
