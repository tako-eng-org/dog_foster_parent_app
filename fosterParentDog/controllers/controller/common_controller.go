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

func strToUint64(i string) uint64 {
	var r, err = strconv.ParseUint(i, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func intToFloat64(i []int) []float64 {
	f := make([]float64, len(i))
	for n := range i {
		f[n] = float64(i[n])
	}
	return f
}

func strToIntSlice(i []string) []int {
	var f []int
	for _, foo := range i {
		j, _ := strconv.Atoi(foo)
		f = append(f, j)
	}
	return f
}
