package storage

import (
	"context"
	"testing"
)

var (
	storage = Storage{Hash: make(map[string]string, 100)}
	ctx     = context.Background()
)

var tests = struct {
	storage Engine
}{
	storage: &Storage{storage.Hash},
}

func TestStorage(t *testing.T) {
	t.Parallel()
	err := tests.storage.Set(ctx, "key", "value")
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = tests.storage.Get(ctx, "key")
	if err != nil {
		t.Fatal(err)
	}

	err = tests.storage.Del(ctx, "key")
	if err != nil {
		t.Fatal(err)
	}
}
