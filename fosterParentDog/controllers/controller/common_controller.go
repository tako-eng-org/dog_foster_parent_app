package controller

import "fpdapp/models/db"

type (
	Controller struct {
		DbConn *db.Database
	}
)
