package compute

import (
	"fmt"
	"strings"
)

const (
	SET = "SET"
	GET = "GET"
	DEL = "DEL"
)

type Query struct {
	Cmd  string
	Args []string
}

func NewQuery(cmd string, args []string) Query {
	return Query{
		Cmd:  cmd,
		Args: args,
	}
}

func Parser(args string) (Query, error) {
	cmd := strings.Fields(args)
	if len(cmd) == 0 {
		return Query{}, nil
	}

	query := Query{}
	switch cmd[0] {
	case SET:
		if len(cmd) != 3 {
			return query, fmt.Errorf("SET expects 2 arguments")
		}
		query.Cmd, query.Args = cmd[0], cmd[1:]
		return query, nil
	case GET:
		if len(cmd) != 2 {
			return query, fmt.Errorf("GET expects 1 arguments")
		}
		query.Cmd, query.Args = cmd[0], cmd[1:]
		return query, nil
	case DEL:
		if len(cmd) != 2 {
			return query, fmt.Errorf("DEL expects 1 arguments")
		}
		query.Cmd, query.Args = cmd[0], cmd[1:]
		return query, nil
	}
	return query, fmt.Errorf("unknown %q command", args)
}
