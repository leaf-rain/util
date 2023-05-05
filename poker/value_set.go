package poker

import "sort"

type valueSet struct {
	value   int64
	times   int64
	isLaizi bool
}

func cardsToValueSet(cards []*Card) ([]*valueSet, int64) {
	var laiziCount int64
	var timesSet []*valueSet
	var isExists bool
	for i := range cards {
		if cards[i].IsLaizi {
			laiziCount++
			continue
		}
		isExists = false
		for _, item := range timesSet {
			if item.value == cards[i].Value {
				isExists = true
				item.times += 1
			}
		}
		if !isExists {
			timesSet = append(timesSet, &valueSet{
				value:   cards[i].Value,
				times:   1,
				isLaizi: cards[i].IsLaizi,
			})
		}
	}
	valueSetSort(timesSet)
	return timesSet, laiziCount
}

func cardsToValueSetOnLaizi(cards []*Card) ([]*valueSet, int64, int64) { // 1:对象，2：癞子数量，3：不是癞子组数
	var laiziCount int64
	var timesSet []*valueSet
	var isExists bool
	var noLaizi int64
	for i := range cards {
		if cards[i].IsLaizi {
			laiziCount++
		}
		isExists = false
		for _, item := range timesSet {
			if item.value == cards[i].Value {
				isExists = true
				item.times += 1
			}
		}
		if !isExists {
			timesSet = append(timesSet, &valueSet{
				value:   cards[i].Value,
				times:   1,
				isLaizi: cards[i].IsLaizi,
			})
			if !cards[i].IsLaizi {
				noLaizi++
			}
		}
	}
	valueSetSortByValue(timesSet)
	return timesSet, laiziCount, noLaizi
}

func valueSetSort(data []*valueSet) {
	sort.Slice(data, func(i, j int) bool {
		if data[i].times == data[j].times {
			return data[i].value > data[j].value
		}
		return data[i].times > data[j].times
	})
}

func valueSetSortByValue(data []*valueSet) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].value > data[j].value
	})
}

func getValueSetByTimes(data []*valueSet) *valueSet {
	if len(data) == 0 {
		return nil
	}
	var result = new(valueSet)
	for _, item := range data {
		if result.times < item.times {
			result = &valueSet{
				value:   item.value,
				times:   item.times,
				isLaizi: item.isLaizi,
			}
		}
	}
	return result
}

func getValueSetByTimesNoJoker(data []*valueSet) *valueSet {
	if len(data) == 0 {
		return nil
	}
	var result = new(valueSet)
	for _, item := range data {
		if result.times < item.times && item.value < littleKing {
			result = &valueSet{
				value:   item.value,
				times:   item.times,
				isLaizi: item.isLaizi,
			}
		}
	}
	return result
}

func getValueSetByValue(data []*valueSet) *valueSet {
	if len(data) == 0 {
		return nil
	}
	valueSetSortByValue(data)
	return data[0]
}

func getValueSetByValueNoJoker(data []*valueSet) *valueSet {
	if len(data) == 0 {
		return nil
	}
	valueSetSortByValue(data)
	for _, item := range data {
		if item.value < littleKing {
			return item
		}
	}
	return nil
}
