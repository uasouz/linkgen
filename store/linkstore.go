package store

// LinkStore Interface describing how a link storage works
type LinkStore interface {
	AddLinkMapping(original, shortID string) error

	GetOriginal(shortID string) (string, error)
}
