package utools

import "time"

// TimeStringToGoTime 字符串转时间
func TimeStringToGoTime(tm string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", tm, time.Local)
	if nil == err && !t.IsZero() {
		return t
	}
	return time.Time{}
}
