package snowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestGetId(t *testing.T) {
	idMap := make(map[int64]struct{})
	var (
		idMax   int64 = 0
		idError int64 = 0
	)
	for i := 0; i < 100; i++ {
		id := GetSnowflakeId().Int64()
		fmt.Println(id)
		fmt.Println(float64(id))
		if id > idMax {
			idMax = id
		} else {
			idError++
		}
		idMap[id] = struct{}{}
	}
	fmt.Println(len(idMap))
	fmt.Println(idError)
	fmt.Println(time.Now().UnixMilli())
}
