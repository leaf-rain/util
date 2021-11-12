package tool

import (
	"math/rand"
	"time"
)

// GetTimeUnix 获取秒级时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// GetTimeUnixMilli 获取毫秒级时间戳
func GetTimeUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetTimeUnixBefore 获取多少时间以前的时间戳
func GetTimeUnixBefore(duration time.Duration) int64 {
	return time.Now().Add(duration).UnixNano() / 1e6
}

// GetRandomRedisTimeOut 在指定时间上添加一个300秒的随机值
func GetRandomRedisTimeOut(base time.Duration) time.Duration {
	rand.Seed(time.Now().UnixNano())
	add := time.Duration(rand.Intn(5*60)) * time.Second
	return base + add
}

// GetEndDaySecond 获取到今天还剩下多少秒
func GetEndDaySecond() int64 {
	t := time.Now()
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return d.AddDate(0, 0, 1).Unix() - GetTimeUnix()
}

// TimeToYYYYMMDDStr 格式化时间
// t 毫秒时间戳
// 1620102515929 -> 2021-05-04
func TimeToYYYYMMDDStr(t int64) string {
	return time.Unix(t/1000, 0).Format("2006-01-02")
}

type DS struct {
	Ds string
	T  time.Time
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []DS {
	var (
		d []DS
	)
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return nil
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return nil
	}
	if date2.Equal(date) {
		d = append(d, DS{
			Ds: date.Format(timeFormatTpl),
			T:  date2,
		})
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return nil
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, DS{
		Ds: date.Format(timeFormatTpl),
		T:  date2,
	})
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, DS{
			Ds: date.Format(timeFormatTpl),
			T:  date,
		})
		if dateStr == date2Str {
			break
		}
	}
	return d
}
