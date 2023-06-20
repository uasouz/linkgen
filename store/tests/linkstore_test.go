package tests

import (
	"linkgen/core"
	"linkgen/store"
	"linkgen/store/memory"
	"testing"
)

var stores = []store.LinkStore{
	memory.New(),
}

func TestLinkStore(t *testing.T) {
	for _, linkStore := range stores {
		shortid, err := core.GenerateNewShortID()
		if err != nil {
			t.Errorf("failed to generate shortid: %v", err)
		}

		err = linkStore.AddLinkMapping("https://google.com", shortid)

		if err != nil {
			t.Errorf("failed to add link mapping: %v", err)
		}

		originalURL, err := linkStore.GetOriginal(shortid)

		if err != nil {
			t.Errorf("failed to get original url: %v", err)
		}

		if originalURL != "https://google.com" {
			t.Errorf("failed to get original url: expected %s, got %s", "https://google.com", originalURL)
		}

		originalURL, err = linkStore.GetOriginal("nonexistent")

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if err != store.ErrLinkNotFound {
			t.Errorf("expected %v, got %v", store.ErrLinkNotFound, err)
		}

		if originalURL != "" {
			t.Errorf("expected empty string, got %s", originalURL)
		}
	}
}
