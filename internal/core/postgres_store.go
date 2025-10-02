package core

import (
	"colonycore/internal/infra/persistence/memory"
	"fmt"
)

// PostgresStore placeholder – embeds an in-memory store so it satisfies the
// PersistentStore interface while real implementation is pending.
type PostgresStore struct{ *memory.Store }

// NewPostgresStore returns a placeholder backed by memory store plus a not-implemented error.
func NewPostgresStore(_ string, engine *RulesEngine) (*PostgresStore, error) {
	ps := &PostgresStore{Store: NewMemoryStore(engine)}
	return ps, fmt.Errorf("postgres driver not yet implemented")
}
