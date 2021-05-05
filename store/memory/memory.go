package memory

type InMemoryLinkStore struct {
	linkMap map[string]string
}

func (m InMemoryLinkStore) AddLinkMapping(original, shortID string) bool {
	panic("implement me")
}

func (m InMemoryLinkStore) GetOriginal(shortID string) string {
	panic("implement me")
}

func New() *InMemoryLinkStore {
	return &InMemoryLinkStore{linkMap: map[string]string{}}
}
