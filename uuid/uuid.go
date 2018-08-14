package uuid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/wuleying/silver-framework/exceptions"
)

// GetUUID
func GetUUID() snowflake.ID {
	node, err := snowflake.NewNode(1)
	exceptions.CheckError(err)
	return node.Generate()
}
