package ws

import "sync"

type Manager struct {
	client *Client

	mu sync.RWMutex
}

var (
	manager *Manager
	once    sync.Once
)

func GetManager() *Manager {

	once.Do(func() {
		manager = &Manager{}
	})

	return manager
}

func (m *Manager) SetClient(c *Client) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.client = c
}

func (m *Manager) Client() *Client {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.client
}

func (m *Manager) Start() error {

	m.mu.RLock()
	client := m.client
	m.mu.RUnlock()

	if client == nil {
		return nil
	}

	return client.Connect()
}

func (m *Manager) Stop() {

	m.mu.RLock()
	client := m.client
	m.mu.RUnlock()

	if client != nil {
		client.Disconnect()
	}
}

func (m *Manager) Subscribe(symbols []string, dataType string) error {

	m.mu.RLock()
	client := m.client
	m.mu.RUnlock()

	if client == nil {
		return nil
	}

	return client.Subscribe(symbols, dataType)
}

func (m *Manager) Status() bool {

	m.mu.RLock()
	client := m.client
	m.mu.RUnlock()

	if client == nil {
		return false
	}

	return client.IsConnected()
}
