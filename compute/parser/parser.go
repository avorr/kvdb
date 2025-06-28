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

func Parser(line string) (Query, error) {
	args := strings.Fields(line)
	if len(args) == 0 {
		return Query{}, nil
	}

	query, cmd, args := Query{}, args[0], args[1:]
	switch cmd {
	case SET:
		if len(args) != 2 {
			return query, fmt.Errorf("SET expects 2 arguments")
		}
		query.Cmd, query.Args = cmd, args
		return query, nil
	case GET, DEL:
		if len(args) != 1 {
			return query, fmt.Errorf("%s expects 1 arguments", cmd)
		}
		query.Cmd, query.Args = cmd, args
		return query, nil
	}
	return query, fmt.Errorf("unknown %q command", args)
}
