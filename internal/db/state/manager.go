package state

import (
	"github.com/NethermindEth/juno/internal/db"
)

// Manager is a database manager, with the objective of managing
// the contract codes and contract storages databases.
type Manager struct {
	stateDatabase db.DatabaseTransactional
	contractDef   db.Database
}

// NewManager returns a new instance of Manager with the given database sources.
func NewManager(stateDatabase db.DatabaseTransactional, contractDef db.Database) *Manager {
	return &Manager{stateDatabase, contractDef}
}

func (m *Manager) Close() {
	m.stateDatabase.Close()
	m.contractDef.Close()
}
