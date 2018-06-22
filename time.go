package got

import "time"

func DayUnix(d int) int64 {
	day := time.Now().AddDate(0, 0, d).Format("20060102")
	t, _ := time.ParseInLocation("20060102", day, time.Local)
	return t.Unix()
}
