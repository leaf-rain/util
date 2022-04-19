package tool

import (
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

var timeTemplates = []string{
	"2006-01-02",
	"2006/01/02",
	"2006-01-02 15:04:05", //常规类型
	//"2006/01/02 15:04:05",
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

func ToYYYYMMDD(t time.Time) int64 {
	return int64((t.Year() * 10000) + (int(t.Month()) * 100) + t.Day())
}

//判断时间是当年的第几周
func WeekByDate(t time.Time) int64 {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int64
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = int64(yearDay-firstWeekDays)/7 + 2
	}
	return week
}
