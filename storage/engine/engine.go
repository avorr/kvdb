package engine

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	compute "kvdb/compute/parser"
	"kvdb/storage"
	"os"
)

type Db struct {
	logger  *zap.Logger
	storage storage.Engine
	query   compute.Query
}

func New(logger *zap.Logger) *Db {
	return &Db{
		logger:  logger,
		storage: storage.New(),
	}
}

func (o *Db) RunQuery() (string, bool, error) {
	switch o.query.Cmd {
	case compute.GET:
		v, ok, err := o.storage.Get(context.Background(), o.query.Args[0])
		return v, ok, err
	case compute.SET:
		err := o.storage.Set(context.Background(), o.query.Args[0], o.query.Args[1])
		return "", false, err
	case compute.DEL:
		err := o.storage.Del(context.Background(), o.query.Args[0])
		return "", false, err
	}
	return "", false, errors.New("unknown command")
}

func (o *Db) Cli() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("cli> ")
	for scanner.Scan() {
		dbQuery, err := compute.Parser(scanner.Text())
		if err != nil {
			fmt.Printf("%s\ncli> ", err)
			continue
		}

		if dbQuery.Args == nil {
			fmt.Print("cli> ")
			continue
		}
		o.query = dbQuery

		v, ok, err := o.RunQuery()
		if err != nil {
			o.logger.Error(err.Error())
		}
		if o.query.Cmd == compute.GET {
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println(nil)
			}
		}
		fmt.Print("cli> ")
	}

	if scanner.Err() != nil {
		o.logger.Fatal(scanner.Err().Error())
	}
}
