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
