package memory

// InMemoryLinkStore An in memory store.LinkStore interface implementation
type InMemoryLinkStore struct {
	linkMap map[string]string
}

// AddLinkMapping - Add an URl to the linkMap with a shortID as key
func (m InMemoryLinkStore) AddLinkMapping(original, shortID string) bool {
	m.linkMap[shortID] = original
	return true
}

// GetOriginal - Retrieve the original URL for the given shortID
func (m InMemoryLinkStore) GetOriginal(shortID string) string {
	return m.linkMap[shortID]
}

// New - creates a new instance for the LinkGen InMemoryLinkStore
func New() *InMemoryLinkStore {
	return &InMemoryLinkStore{linkMap: map[string]string{}}
}
