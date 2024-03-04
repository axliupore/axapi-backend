package utils

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析字符串表示的时间间隔，并返回 time.Duration 对象
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	if strings.Contains(d, "d") {
		// 如果包含 "d" 后缀，解析天数
		index := strings.Index(d, "d")
		days, err := strconv.Atoi(d[:index])
		if err != nil {
			return 0, err
		}

		// 将天数转换为小时，然后创建时间间隔
		duration := time.Duration(days) * 24 * time.Hour
		return duration, nil
	}

	// 如果不包含 "d" 后缀，尝试解析为整数，表示天数
	days, err := strconv.Atoi(d)
	if err != nil {
		return 0, err
	}

	// 将天数转换为小时，然后创建时间间隔
	duration := time.Duration(days) * 24 * time.Hour
	return duration, nil
}
