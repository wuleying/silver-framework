package uuid

import (
	"github.com/bwmarrin/snowflake"
)

// GetUUID
func GetUUID() (snowflake.ID, error) {
	var UUID snowflake.ID

	node, err := snowflake.NewNode(1)
	if err != nil {
		return UUID, err
	}

	UUID = node.Generate()

	return UUID, nil
}
