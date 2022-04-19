package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

type ttt struct {
	A string
	B string
	C float64
}

func main() {
	var f float64 = 100.0014
	fmt.Println(*(*uint64)(unsafe.Pointer(&f)))
	fmt.Println(int64(f))
}

//判断时间是当年的第几周
func WeekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("%d第%d周", t.Year(), week)
}

type WeekDate struct {
	WeekTh    string
	StartTime time.Time
	EndTime   time.Time
}

// 将开始时间和结束时间分割为周为单位
func GroupByWeekDate(startTime, endTime time.Time) []WeekDate {
	weekDate := make([]WeekDate, 0)
	diffDuration := endTime.Sub(startTime)
	days := int(math.Ceil(float64(diffDuration/(time.Hour*24)))) + 1

	currentWeekDate := WeekDate{}
	currentWeekDate.WeekTh = WeekByDate(endTime)
	currentWeekDate.EndTime = endTime
	currentWeekDay := int(endTime.Weekday())
	if currentWeekDay == 0 {
		currentWeekDay = 7
	}
	currentWeekDate.StartTime = endTime.AddDate(0, 0, -currentWeekDay+1)
	nextWeekEndTime := currentWeekDate.StartTime
	weekDate = append(weekDate, currentWeekDate)

	for i := 0; i < (days-currentWeekDay)/7; i++ {
		weekData := WeekDate{}
		weekData.EndTime = nextWeekEndTime
		weekData.StartTime = nextWeekEndTime.AddDate(0, 0, -7)
		weekData.WeekTh = WeekByDate(weekData.StartTime)
		nextWeekEndTime = weekData.StartTime
		weekDate = append(weekDate, weekData)
	}

	if lastDays := (days - currentWeekDay) % 7; lastDays > 0 {
		lastData := WeekDate{}
		lastData.EndTime = nextWeekEndTime
		lastData.StartTime = nextWeekEndTime.AddDate(0, 0, -lastDays)
		lastData.WeekTh = WeekByDate(lastData.StartTime)
		weekDate = append(weekDate, lastData)
	}

	return weekDate
}

func CheckRepeatForList(l1, l2 []uint64) ([]uint64, float64) {
	var resultList = make([]uint64, 0)
	var resultNum float64
	var length1 = len(l1)
	var length2 = len(l2)
	if length1 <= 1024 || length2 <= 1024 { // golang中小于1024的话，直接下标遍历数组，会优于map
		for i1 := range l1 {
			for i2 := range l2 {
				if l1[i1] == l2[i2] {
					resultList = append(resultList, l1[i1])
					resultNum += 1
				}
			}
		}
		return resultList, resultNum
	}
	var m = make(map[uint64]struct{})
	if length1 > length2 {
		for i := range l1 {
			m[l1[i]] = struct{}{}
		}
		for i := range l2 {
			if _, ok := m[l2[i]]; ok {
				resultList = append(resultList, l2[i])
				resultNum += 1
			}
		}
		return resultList, resultNum
	} else {
		for i := range l2 {
			m[l2[i]] = struct{}{}
		}
		for i := range l1 {
			if _, ok := m[l1[i]]; ok {
				resultList = append(resultList, l1[i])
				resultNum += 1
			}
		}
		return resultList, resultNum
	}
}
