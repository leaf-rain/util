package poker

import "testing"

func Test_valueSetSort(t *testing.T) {
	var req = []*valueSet{
		{
			value: 1,
			times: 3,
		},
		{
			value: 2,
			times: 3,
		},
		{
			value: 3,
			times: 1,
		},
	}
	valueSetSort(req)
	for _, item := range req {
		t.Log(item)
	}
}
