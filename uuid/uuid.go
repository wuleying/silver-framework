package uuid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/exceptions"
)

// GetUUID
func GetUUID() snowflake.ID {
	node, err := snowflake.NewNode(1)
	exceptions.CheckError(err)

	UUID := node.Generate()
	clog.Info("Request_id: %s", UUID.Base58())

	return UUID
}
