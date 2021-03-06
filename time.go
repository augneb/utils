package util

import (
	"time"
	"strings"
)

func Date(format string, ts ...time.Time) string {
	patterns := []string{
		"Y", "2006", // 4 位数字完整表示的年份
		"m", "01",   // 数字表示的月份，有前导零
		"d", "02",   // 月份中的第几天，有前导零的 2 位数字
		"H", "15",   // 小时，24 小时格式，有前导零
		"i", "04",   // 有前导零的分钟数
		"s", "05",   // 秒数，有前导零
	}

	format = strings.NewReplacer(patterns...).Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}

	return t.Format(format)
}

func StrToTime(value string) (time.Time, error) {
	layouts := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006-01-02",
		"2006/01/02",
		"20060102150405",
	}

	var t time.Time
	var e error
	for _, layout := range layouts {
		t, e = time.ParseInLocation(layout, value, time.Local)
		if e == nil {
			return t, nil
		}
	}

	return t, e
}

func Timestamp(v ...interface{}) int64 {
	b := 10
	if len(v) > 0 {
		b = v[0].(int)
	}

	var n time.Time
	if len(v) > 1 {
		n = v[1].(time.Time)
	} else {
		n = time.Now()
	}

	switch b {
	case 10:
		return n.UnixNano() / int64(time.Second)
	case 13:
		return n.UnixNano() / int64(time.Millisecond)
	default:
		return n.UnixNano() / int64(time.Millisecond)
	}
}
