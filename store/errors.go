package store

const (
	ErrLinkNotFound = Error("Link not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
