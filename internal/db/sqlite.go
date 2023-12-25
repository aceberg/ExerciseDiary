package db

import (
	"sync"

	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

var mu sync.Mutex

func connect(path string) *sqlx.DB {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	return dbx
}

func exec(path string, sqlStatement string) {

	mu.Lock()
	dbx := connect(path)
	_, err := dbx.Exec(sqlStatement)
	mu.Unlock()

	check.IfError(err)
}

// SelectEx - select all exercises from DB
func SelectEx(path string) (exes []models.Exercise) {

	mu.Lock()
	dbx := connect(path)
	err := dbx.Select(&exes, "SELECT * FROM exercises ORDER BY ID ASC")
	mu.Unlock()

	check.IfError(err)

	return exes
}

// SelectSet - select all sets from DB
func SelectSet(path string) (sets []models.Set) {

	mu.Lock()
	dbx := connect(path)
	err := dbx.Select(&sets, "SELECT * FROM sets ORDER BY ID ASC")
	mu.Unlock()

	check.IfError(err)

	return sets
}

// SelectW - select all weight from DB
func SelectW(path string) (w []models.BodyWeight) {

	mu.Lock()
	dbx := connect(path)
	err := dbx.Select(&w, "SELECT * FROM weight ORDER BY ID ASC")
	mu.Unlock()

	check.IfError(err)

	return w
}
