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
	data map[string]string
}

func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (o *Storage) Set(ctx context.Context, key, value string) error {
	o.data[key] = value
	return nil
}

func (o *Storage) Get(ctx context.Context, key string) (string, bool, error) {
	v, ok := o.data[key]
	return v, ok, nil
}

func (o *Storage) Del(ctx context.Context, key string) error {
	delete(o.data, key)
	return nil
}
