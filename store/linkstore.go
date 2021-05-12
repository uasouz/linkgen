package store

// LinkStore Interface describing how a link storage works
type LinkStore interface {
	AddLinkMapping(original, shortID string) bool

	GetOriginal(shortID string) string
}
