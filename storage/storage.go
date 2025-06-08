package storage

import (
	"context"
)

type Engine interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, bool, error)
	Del(context.Context, string) error
}

type Storage struct {
	Hash map[string]string
}

func New() *Storage {
	return &Storage{
		Hash: make(map[string]string, 100),
	}
}

func (o *Storage) Set(ctx context.Context, key, value string) error {
	o.Hash[key] = value
	return nil
}

func (o *Storage) Get(ctx context.Context, key string) (string, bool, error) {
	v, ok := o.Hash[key]
	return v, ok, nil
}

func (o *Storage) Del(ctx context.Context, key string) error {
	delete(o.Hash, key)
	return nil
}
