package plugins

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

type Manager struct {
	engine  *gin.Engine
	plugins map[string]Plugin
	mu      sync.RWMutex
}

func NewManager(engine *gin.Engine) *Manager {
	return &Manager{
		engine:  engine,
		plugins: make(map[string]Plugin),
	}
}

func (m *Manager) Install(p Plugin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[p.Name()]; exists {
		return fmt.Errorf("plugin %s already installed", p.Name())
	}

	if err := p.Init(); err != nil {
		return fmt.Errorf("plugin %s init failed: %v", p.Name(), err)
	}

	m.plugins[p.Name()] = p

	// 注册路由
	if router := p.Router(); router != nil {
		router(m.engine)
	}

	return nil
}

func (m *Manager) Uninstall(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.plugins[name]
	if !exists {
		return fmt.Errorf("plugin %s not found", name)
	}

	if err := p.Destroy(); err != nil {
		return fmt.Errorf("plugin %s destroy failed: %v", name, err)
	}

	delete(m.plugins, name)
	return nil
}

func (m *Manager) GetPlugin(name string) (Plugin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	p, exists := m.plugins[name]
	return p, exists
}

func (m *Manager) ListPlugins() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.plugins))
	for name := range m.plugins {
		names = append(names, name)
	}
	return names
}
