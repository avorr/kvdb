package engine

import (
	"context"
	"go.uber.org/zap"
	"log"
	"testing"
)

var (
	ctx = context.Background()
)

func TestEngine(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	db := New(logger)
	t.Parallel()
	err = db.storage.Set(ctx, "key", "value")
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
