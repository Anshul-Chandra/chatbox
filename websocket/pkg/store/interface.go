package store

import (
	"github.com/chatbox/websocket/pkg/types"
)

// Store is an interface for the system that maintains list of client connections
type Store interface {
	Register(c *types.Connection) error
	Unregister(connectionID string) error
	Get(connectionID string) ([]*types.Connection, error)
}
