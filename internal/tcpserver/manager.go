package tcpserver

import "sync"

// manage all tcp connections
type Manager struct {
	sync.Map
}

var manager Manager

func (m *Manager) Store(id int64, ctx *ConnCtx) {
	m.Map.Store(id, ctx)
}

func (m *Manager) Load(id int64) *ConnCtx {
	value, ok := m.Map.Load(id)
	if !ok {
		return nil
	}
	return value.(*ConnCtx)
}

func (m *Manager) Delete(id int64) {
	m.Map.Delete(id)
}
