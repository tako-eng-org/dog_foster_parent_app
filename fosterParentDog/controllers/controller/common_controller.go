package controller

import "fpdapp/models/db"

type (
	Controller struct {
		Database *db.Database
	}
)
