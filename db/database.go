package main

import "github.com/go-pg/pg/v10"

func PostgresConnect() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "pets",
	}
	return pg.Connect(options)
}
