package engine

import (
	"context"
	"errors"
	"go.uber.org/zap"
	compute "kvdb/compute/parser"
	"kvdb/storage"
	"log"
)

type Db struct {
	Logger  *zap.Logger
	storage storage.Engine
	compute.Query
}

func New() *Db {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	return &Db{
		Logger:  logger,
		storage: storage.New(),
	}
}

func (o *Db) RunQuery() (string, bool, error) {
	switch o.Cmd {
	case compute.GET:
		v, ok, err := o.storage.Get(context.Background(), o.Args[0])
		return v, ok, err
	case compute.SET:
		err := o.storage.Set(context.Background(), o.Args[0], o.Args[1])
		return "", false, err
	case compute.DEL:
		err := o.storage.Del(context.Background(), o.Args[0])
		return "", false, err
	}
	return "", false, errors.New("unknown command")
}
