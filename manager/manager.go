package manager

import "context"

func NewManager() *Manager {
	m := &Manager{

	}
	return m
}

type Manager struct {
}

func (m *Manager) Run(ctx context.Context) {

}
