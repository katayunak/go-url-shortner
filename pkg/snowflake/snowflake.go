package snowflake

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func Init(nodeID int64) error {
	newNode, err := snowflake.NewNode(nodeID)
	if err != nil {
		return err
	}

	node = newNode
	return nil
}

func GenerateID() int64 {
	return node.Generate().Int64()
}
