package snowflake

import (
	"fmt"

	sf "github.com/bwmarrin/snowflake"
)

type Generator struct {
	node *sf.Node
}

func (sg *Generator) GenerateID() int64 {
	return sg.node.Generate().Int64()
}

func New(nodeId int64) (*Generator, error) {
	node, err := sf.NewNode(nodeId)

	if err != nil {
		return nil, fmt.Errorf("create snowflake node: %w", err)
	}

	return &Generator{
		node: node,
	}, nil
}
