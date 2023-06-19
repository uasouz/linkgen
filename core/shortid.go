package core

import nanoid "github.com/matoous/go-nanoid/v2"

// GenerateNewShortID - generates a new ShortID using the implemented method, currently using gonanoid
func GenerateNewShortID() (string, error) {
	return nanoid.New(8)
}
