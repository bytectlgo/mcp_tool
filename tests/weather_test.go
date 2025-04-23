package weather_test

import (
	"testing"

	"github.com/bytectlgo/mcp_tool/pkg/weather"
)

func TestParseQuery(t *testing.T) {
	tests := []struct {
		query string
		want  weather.QueryParams
	}{
		{
			"北京的天气",
			weather.QueryParams{City: "北京", Days: 1, Units: "metric", Lang: "zh"},
		},
		{
			"上海未来3天天气",
			weather.QueryParams{City: "上海", Days: 3, Units: "metric", Lang: "zh"},
		},
		{
			"伦敦明天会下雨吗",
			weather.QueryParams{City: "伦敦", Days: 1, Units: "metric", Lang: "zh"},
		},
		{
			"纽约的天气 fahrenheit",
			weather.QueryParams{City: "纽约", Days: 1, Units: "imperial", Lang: "zh"},
		},
		{
			"东京天气 english",
			weather.QueryParams{City: "东京", Days: 1, Units: "metric", Lang: "en"},
		},
	}

	for _, test := range tests {
		got := weather.ParseQuery(test.query)
		if got != test.want {
			t.Errorf("ParseQuery(%q) = %+v, want %+v", test.query, got, test.want)
		}
	}
}
