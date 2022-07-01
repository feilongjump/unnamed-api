package app

import (
	"time"
	"unnamed-api/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}
