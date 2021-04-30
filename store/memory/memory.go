package memory

type LinkStore struct {
	linkMap map[string]string
}

func (m LinkStore) AddLinkMapping(original, shortID string) bool {
	panic("implement me")
}

func (m LinkStore) GetOriginal(shortID string) string {
	panic("implement me")
}

func New() *LinkStore {
	return &LinkStore{linkMap: map[string]string{}}
}
