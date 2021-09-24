package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

const (
	maxNode = 1023
)

func GetSnowflakeId() snowflake.ID {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(maxNode)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return node.Generate()
}
func SnowflakeBase2() string {
	return GetSnowflakeId().Base2()
}
func SnowflakeBase32() string {
	return GetSnowflakeId().Base32()
}
func SnowflakeBase36() string {
	return GetSnowflakeId().Base36()
}
func SnowflakeBase58() string {
	return GetSnowflakeId().Base58()
}
func SnowflakeBase64() string {
	return GetSnowflakeId().Base64()
}
func SnowflakeString() string {
	return GetSnowflakeId().String()
}
func SnowflakeBytes() []byte {
	return GetSnowflakeId().Bytes()
}
func SnowflakeInt64() int64 {
	return GetSnowflakeId().Int64()
}
func SnowflakeIntBytes() [8]byte {
	return GetSnowflakeId().IntBytes()
}
