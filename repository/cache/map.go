package cache

import (
	"context"
	"errors"
	"time"
)

func (m *MapCache) Get(ctx context.Context, key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.mapCache[key]
	if !ok {
		return "", errors.New("key not found")
	}

	return value, nil
}

func (m *MapCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.mapCache[key] = value
	return nil
}
