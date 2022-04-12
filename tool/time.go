package tool

import (
	"math/rand"
	"time"
)

var (
	IST, _ = time.LoadLocation("Asia/Kolkata")
)

// 获取当前天n天钱0点时间戳和当天24点时间戳, 传0即为当前天的0点和24点
func GetBeforeToDayTimeStampRangeByDay(day int) (int64, int64) {
	now := time.Now()
	nowStamp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	before := nowStamp.AddDate(0, 0, 0-day)
	return before.Unix(), nowStamp.Unix() + 86400
}

// 获取当前天的0点和24点
func GetToDayTimeStampRangeByDay(t time.Time) (int64, int64) {
	stamp := GetZeroTime(t).Unix()
	return stamp, stamp + 86400
}

// 获取t周第一天0点时间戳和本周最后一天24点时间戳
func GetToDayTimeStampRangeByWeek(t time.Time) (int64, int64) {
	stamp := GetFirstDateOfWeek(t).Unix()
	return stamp, stamp + 604800
}

// 获取t月第一天0点时间戳和本周最后一天24点时间戳
func GetToDayTimeStampRangeByMonth(t time.Time) (int64, int64) {
	return GetFirstDateOfMonth(t).Unix(), GetLastDateOfMonth(t).Unix()
}

// 获取当天0点
func GetZeroTimeStampToDay() int64 {
	now := time.Now()
	nowStamp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return nowStamp.Unix()
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取本周周一时间
func GetFirstDateOfWeek(now time.Time) time.Time {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

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

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	//"2006/01/02 15:04:05",
	"2006-01-02",
	//"2006/01/02",
}

func GetTimeStampByString(str string, loc *time.Location) int64 {
	var t time.Time
	var err error
	for i := range timeTemplates {
		t, err = time.ParseInLocation(timeTemplates[i], str, loc)
		if err == nil {
			return t.Unix()
		}
	}
	return 0
}
