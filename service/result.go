package service

import (
	"dfc/db"
)

type Result struct {
	ID uint `json:"id"`
}

type ResultServices struct {
	Result      Result
	ResultStore db.Store
}

func NewResultServices(t Result, tStore db.Store) *ResultServices {
	return &ResultServices{
		Result:      t,
		ResultStore: tStore,
	}
}
