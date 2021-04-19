package controller

import (
	"fpdapp/models/db"
	"log"
	"strconv"
)

type (
	Controller struct {
		DbConn *db.Database
	}
)

func strToInt(i string) int {
	r, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func strToInt64(i string) uint64 {
	var r, err = strconv.ParseUint(i, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
