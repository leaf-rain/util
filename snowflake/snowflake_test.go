package snowflake

import (
	"fmt"
	"testing"
)

func TestGetId(t *testing.T) {
	id := SnowflakeId()
	fmt.Println(id)
}
