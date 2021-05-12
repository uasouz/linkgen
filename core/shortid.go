package core

import nanoid "github.com/matoous/go-nanoid/v2"

func GenerateNewShortID() (string, error) {
	return nanoid.New(8)
}
