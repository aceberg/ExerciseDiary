package db

import (
	"github.com/aceberg/ExerciseDiary/internal/models"
)

// BulkAddSets - add slice
func BulkAddSets(path string, allExs []models.Set) {
	var oneEx models.Set

	for _, oneEx = range allExs {
		InsertSet(path, oneEx)
	}
}

// BulkDeleteSetsByDate - delete all Sets with date
func BulkDeleteSetsByDate(path, date string) {
	var oneEx models.Set

	allExs := SelectSet(path)

	for _, oneEx = range allExs {
		if oneEx.Date == date {
			DeleteSet(path, oneEx.ID)
		}
	}
}
