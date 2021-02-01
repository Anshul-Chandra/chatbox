package store

import (
	"errors"

	"github.com/chatbox/websocket/pkg/types"
)

// InMemoryStore implements Store
type InMemoryStore struct {
	connections map[string]*types.Connection
}

// NewInMemoryStore returns a new instance of InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		connections: make(map[string]*types.Connection),
	}
}

// Register adds a new client in the store
func (store *InMemoryStore) Register(conn *types.Connection) error {
	if store == nil {
		return errors.New("store object is nil")
	}

	if store.connections == nil {
		store.connections = make(map[string]*types.Connection)
	}

	id := conn.GetID()
	store.connections[id] = conn
	return nil
}

// Unregister removes a given client from the store
func (store *InMemoryStore) Unregister(connectionID string) error {
	if store == nil {
		return errors.New("store object is nil")
	}

	if _, present := store.connections[connectionID]; !present {
		return errors.New("invalid client")
	}

	delete(store.connections, connectionID)
	return nil
}

// Get returns the client connection(s) corresponding to a connection ID
func (store *InMemoryStore) Get(connectionID string) ([]*types.Connection, error) {
	if store == nil {
		return nil, errors.New("store has not been initialized yet")
	}

	conn, present := store.connections[connectionID]
	if !present {
		return nil, errors.New("invalid client id")
	}

	return []*types.Connection{conn}, nil
}
