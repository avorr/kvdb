package engine

import (
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestEngine(t *testing.T) {
	db := New()
	t.Parallel()
	err := db.storage.Set(ctx, "key", "value")
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = db.storage.Get(ctx, "key")
	if err != nil {
		t.Fatal(err)
	}

	err = db.storage.Del(ctx, "key")
	if err != nil {
		t.Fatal(err)
	}
}
