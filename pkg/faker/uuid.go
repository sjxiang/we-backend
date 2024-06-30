package faker

import (
	"time"
	
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.NewString()
}

func ValidateUUID(value string) bool {
	if _, err := uuid.Parse(value); err != nil {
		return false
	}

	return true
}


type idGenerator struct {
	node *snowflake.Node
}

func NewIDGenerator(now time.Time, machineId int64) (*idGenerator, error) {

	snowflake.Epoch = now.Unix()
	node, err := snowflake.NewNode(machineId)
	if err != nil {
		return nil, err
	}

	return &idGenerator{node: node}, nil
}

func (i *idGenerator) GenerateUUID() int64 {
	return i.node.Generate().Int64()
}
