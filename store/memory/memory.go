package memory

import "linkgen/store"

// InMemoryLinkStore An in memory store.LinkStore interface implementation
type InMemoryLinkStore struct {
	linkMap map[string]string
}

// AddLinkMapping - Add an URl to the linkMap with a shortID as key
func (m InMemoryLinkStore) AddLinkMapping(original, shortID string) error {
	m.linkMap[shortID] = original
	return nil
}

// GetOriginal - Retrieve the original URL for the given shortID
func (m InMemoryLinkStore) GetOriginal(shortID string) (string, error) {
	if _, ok := m.linkMap[shortID]; !ok {
		return "", store.ErrLinkNotFound
	}
	return m.linkMap[shortID], nil
}

// New - creates a new instance for the LinkGen InMemoryLinkStore
func New() *InMemoryLinkStore {
	return &InMemoryLinkStore{linkMap: map[string]string{}}
}
