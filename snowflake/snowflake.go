package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func SnowflakeId() string {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Generate a snowflake ID.
	id := node.Generate()
	return id.Base58()
}
