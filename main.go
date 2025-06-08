package main

import (
	"bufio"
	"fmt"
	parser "kvdb/compute/parser"
	"kvdb/storage/engine"
	"os"
)

func main() {
	db := engine.New()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("cli> ")
	for scanner.Scan() {
		dbQuery, err := parser.Parser(scanner.Text())
		db.Query = dbQuery
		if err != nil {
			fmt.Printf("%s\ncli> ", err)
			continue
		}

		v, ok, err := db.RunQuery()
		if err != nil {
			db.Logger.Error(err.Error())
		}
		if db.Cmd == parser.GET {
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println(nil)
			}
		}
		fmt.Print("cli> ")
	}

	if scanner.Err() != nil {
		db.Logger.Fatal(scanner.Err().Error())
	}
}
