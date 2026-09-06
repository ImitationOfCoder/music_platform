package snowflake

import (
	"fmt"

	sf "github.com/bwmarrin/snowflake"
	"github.com/kelseyhightower/envconfig"
)

type SnowflakeGenerator struct {
	node *sf.Node
}

type Config struct {
	NodeID int64 `envconfig:"SNOWFLAKE_NODE_ID" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process config: %w", err)
	}

	return config, nil
}

func NewSnowflakeGenerator() (*SnowflakeGenerator, error) {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get snowflake config: %w", err)
		panic(err)
	}

	node, err := sf.NewNode(config.NodeID)

	if err != nil {
		return nil, fmt.Errorf("create snowflake node: %w", err)
	}

	return &SnowflakeGenerator{
		node: node,
	}, nil
}

func (sg *SnowflakeGenerator) GenerateID() int64 {
	return sg.node.Generate().Int64()
}
