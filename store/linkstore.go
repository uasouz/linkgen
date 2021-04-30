package store

type LinkStore interface {
	AddLinkMapping(original, shortID string) bool

	GetOriginal(shortID string) string
}